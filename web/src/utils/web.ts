import utils from "./utils"
/**
 * 页面专用
 */
window.systemInfo = window.systemInfo || {}
const webTs = {
    /**
     * 接口地址
     * @param str
     * @returns {string|string|*}
     */
    apiUrl(str) {
        if (str == "privacy") {
            const apiHome = utils.getDomain(window.systemInfo.apiUrl)
            if (apiHome == "" || apiHome == "public") {
                return "https://www.dootask.com/privacy.html"
            }
            str = "../privacy.html"
        }
        if (
            str.substring(0, 2) === "//" ||
            str.substring(0, 7) === "http://" ||
            str.substring(0, 8) === "https://" ||
            str.substring(0, 6) === "ftp://" ||
            str.substring(0, 1) === "/"
        ) {
            return str
        }
        if (typeof window.systemInfo.apiUrl === "string") {
            str = window.systemInfo.apiUrl + str
        } else {
            str = window.location.origin + "/api/" + str
        }
        while (str.indexOf("/../") !== -1) {
            str = str.replace(/\/(((?!\/).)*)\/\.\.\//, "/")
        }
        return str
    },

    /**
     * 返回对话未读数量（不含免打扰，但如果免打扰中有@则返回@数量）
     * @param dialog
     * @returns {*|number}
     */
    getDialogNum(dialog) {
        if (!dialog) {
            return 0
        }
        const unread = !dialog.silence ? dialog.unread : 0
        return unread || dialog.mention || dialog.mark_unread || 0
    },

    /**
     * 返回对话未读数量
     * @param dialog
     * @param containSilence    是否包含免打扰消息（true:包含, false:不包含）
     * @returns {*|number}
     */
    getDialogUnread(dialog, containSilence) {
        if (!dialog) {
            return 0
        }
        const unread = (containSilence || !dialog.silence) ? dialog.unread : 0
        return unread || dialog.mark_unread || 0
    },
    /**
     * 对话标签
     * @param dialog
     * @returns {*[]}
     */
    dialogTags(dialog) {
        let tags = [];
        if (dialog.type == 'group') {
            if (['project', 'task'].includes(dialog.group_type) && utils.isJson(dialog.group_info)) {
                if (dialog.group_type == 'task' && dialog.group_info.complete_at) {
                    tags.push({
                        color: 'success',
                        text: $t('已完成')
                    })
                }
                if (dialog.group_info.deleted_at) {
                    tags.push({
                        color: 'red',
                        text: $t('已删除')
                    })
                } else if (dialog.group_info.archived_at) {
                    tags.push({
                        color: 'default',
                        text: $t('已归档')
                    })
                }
            }
        }
        return tags;
    },

    /**
 * 对话完成
 * @param dialog
 * @returns {*[]}
 */
    dialogCompleted(dialog) {
        return webTs.dialogTags(dialog).find(({ color }) => color == 'success');
    },

    /**
 * 返回文本信息预览格式
 * @param text
 * @returns {*}
 */
    getMsgTextPreview(text) {
        if (!text) return '';
        text = text.replace(/<img\s+class="emoticon"[^>]*?alt="(\S+)"[^>]*?>/g, "[$1]")
        text = text.replace(/<img\s+class="emoticon"[^>]*?>/g, `[${$t('动画表情')}]`)
        text = text.replace(/<img\s+class="browse"[^>]*?>/g, `[${$t('图片')}]`)
        text = text.replace(/<[^>]+>/g, "")
        text = text.replace(/&nbsp;/g, " ")
        text = text.replace(/&amp;/g, "&")
        text = text.replace(/&lt;/g, "<")
        text = text.replace(/&gt;/g, ">")
        return text
    },
    /**
     * 消息简单描述
     * @param data
     * @returns {string|*}
     */
    getMsgSimpleDesc(data) {
        if (utils.isJson(data)) {
            switch (data.type) {
                case 'text':
                    return webTs.getMsgTextPreview(data.msg.text)
                case 'record':
                    return `[${$t('语音')}]`
                case 'meeting':
                    return `[${$t('会议')}] ${data.msg.name}`
                case 'file':
                    if (data.msg.type == 'img') {
                        return `[${$t('图片')}]`
                    }
                    return `[${$t('文件')}] ${data.msg.name}`
                case 'tag':
                    return `[${$t(data.msg.action === 'remove' ? '取消标注' : '标注')}] ${webTs.getMsgSimpleDesc(data.msg.data)}`
                case 'todo':
                    return `[${$t(data.msg.action === 'remove' ? '取消待办' : (data.msg.action === 'done' ? '完成' : '设待办'))}] ${webTs.getMsgSimpleDesc(data.msg.data)}`
                case 'notice':
                    return data.msg.notice
                default:
                    return `[${$t('未知的消息')}]`
            }
        }
        return '';
    },

    /**
        * 消息格式化处理
        * @param text
        * @param userid
        * @returns {string|*}
        */
    formatTextMsg(text, userid) {
        if (!text) {
            return ""
        }
        const atReg = new RegExp(`<span class="mention user" data-id="${userid}">`, "g")
        text = text.trim().replace(/(\n\x20*){3,}/g, "\n\n");
        text = text.replace(/&nbsp;/g, ' ')
        text = text.replace(/<p><\/p>/g, '<p><br/></p>')
        text = text.replace(/\{\{RemoteURL\}\}/g, webTs.apiUrl('../'))
        text = text.replace(atReg, `<span class="mention me" data-id="${userid}">`)
        // 处理内容连接
        if (/https*:\/\//.test(text)) {
            text = text.split(/(<[^>]*>)/g).map(string => {
                if (string && !/<[^>]*>/.test(string)) {
                    string = string.replace(/(^|[^'"])((https*:\/\/)((\w|=|\?|\.|\/|&|-|:|\+|%|;|#|@|,|!)+))/g, "$1<a href=\"$2\" target=\"_blank\">$2</a>")
                }
                return string;
            }).join("")
        }
        // 处理图片显示尺寸
        const array = text.match(/<img\s+[^>]*?>/g);
        if (array) {
            const widthReg = new RegExp("width=\"(\\d+)\""),
                heightReg = new RegExp("height=\"(\\d+)\"")
            array.some(res => {
                const widthMatch = res.match(widthReg),
                    heightMatch = res.match(heightReg);
                if (widthMatch && heightMatch) {
                    const width = parseInt(widthMatch[1]),
                        height = parseInt(heightMatch[1]),
                        maxSize = res.indexOf("emoticon") > -1 ? 150 : 220;
                    const scale = utils.scaleToScale(width, height, maxSize, maxSize);
                    const value = res
                        .replace(widthReg, `original-width="${width}" width="${scale.width}"`)
                        .replace(heightReg, `original-height="${height}" height="${scale.height}"`)
                    text = text.replace(res, value)
                } else {
                    text = text.replace(res, `<div class="no-size-image-box">${res}</div>`);
                }
            })
        }
        return text;
    },
    /**
     * 格式化时间
     * @param date
     * @returns {*|string}
     */
    formatTime(date) {
        let now = utils.Time(),
            time = Number(utils.Date(date, true)),
            string = '';
        if (Math.abs(now - time) < 3600 * 6 || utils.formatDate('Ymd', now) === utils.formatDate('Ymd', time)) {
            string = utils.formatDate('H:i', time)
        } else if (utils.formatDate('Y', now) === utils.formatDate('Y', time)) {
            string = utils.formatDate('m-d', time)
        } else {
            string = utils.formatDate('Y-m-d', time)
        }
        return string || '';
    },

    /**
 * 服务地址
 * @param str
 * @returns {string}
 */
    originUrl(str) {
        if (str.substring(0, 2) === "//" ||
            str.substring(0, 7) === "http://" ||
            str.substring(0, 8) === "https://" ||
            str.substring(0, 6) === "ftp://" ||
            str.substring(0, 1) === "/") {
            return str;
        }
        if (typeof window.systemInfo.origin === "string") {
            str = window.systemInfo.origin + str;
        } else {
            str = window.location.origin + "/" + str;
        }
        while (str.indexOf("/../") !== -1) {
            str = str.replace(/\/(((?!\/).)*)\/\.\.\//, "/")
        }
        return str
    },

    /**
     * 小于9补0
     * @param val
     * @returns {number|string}
     */
    formatBit(val) {
        val = +val
        return val > 9 ? val : '0' + val
    },

    /**
     * 秒转时间
     * @param second
     * @returns {string}
     */
    formatSeconds(second) {
        let duration
        let days = Math.floor(second / 86400);
        let hours = Math.floor((second % 86400) / 3600);
        let minutes = Math.floor(((second % 86400) % 3600) / 60);
        let seconds = Math.floor(((second % 86400) % 3600) % 60);
        if (days > 0) {
            if (hours > 0) duration = days + "d," + webTs.formatBit(hours) + "h";
            else if (minutes > 0) duration = days + "d," + webTs.formatBit(minutes) + "min";
            else if (seconds > 0) duration = days + "d," + webTs.formatBit(seconds) + "s";
            else duration = days + "d";
        }
        else if (hours > 0) duration = webTs.formatBit(hours) + ":" + webTs.formatBit(minutes) + ":" + webTs.formatBit(seconds);
        else if (minutes > 0) duration = webTs.formatBit(minutes) + ":" + webTs.formatBit(seconds);
        else if (seconds > 0) duration = webTs.formatBit(seconds) + "s";
        return duration;
    },

    /**
     * 倒计时格式
     * @param date
     * @param nowTime
     * @returns {string|*}
     */
    countDownFormat(date, nowTime) {
        let time = Math.round(new Date(date).getTime() / 1000) - nowTime;
        if (time < 86400 * 7 && time > 0) {
            return webTs.formatSeconds(time);
        } else if (time < 0) {
            return '-' + webTs.formatSeconds(time * -1);
        } else if (time == 0) {
            return 0 + 's';
        }
        return webTs.formatTime(date)
    },
}
export default webTs
