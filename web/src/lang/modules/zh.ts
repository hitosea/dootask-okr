const message = {
    commons: {
        http: {
            dataFormatError: '返回数据格式错误',
            reqFailed: '请求失败',
        },
        dialog: {
            title: '提示',
            warmTitle: '温馨提示',
            errTitle: '错误提示',
            successTitle: '成功提示',
            info: '详细信息',
            okText: '确定',
            cancelText: '取消',
        }
    },
    login: {
        reg: '注册',
        login: '登录',
        email: '邮箱',
        enterEmail: '请输入邮箱地址',
        password: '密码',
        enterPassword: '请输入密码',
        password2: '确认密码',
        enterPassword2: '请输入确认密码',
        agree: '登录即表示您同意我们的服务条款和隐私政策。',
        emailError: '邮箱格式错误',
        emailEmpty: '邮箱不能为空',
        passwordLengthError: '密码长度必须大于6位小于20位',
        passwordEmpty: '密码不能为空',
        passwordDiff: '两次密码不一致',
        password2Empty: '确认密码不能为空',
    },
    result: {
        home: '返回首页',
    },
    main: {
        welcome: '您好, 欢迎使用！',
    },
};
export default {
    ...message,
};
