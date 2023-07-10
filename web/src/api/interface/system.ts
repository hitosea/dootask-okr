/* eslint-disable @typescript-eslint/no-namespace */
export namespace System {
    // 系统设置设置
    export interface systemSetting {
        all_group_autoin?: any
        all_group_mute?: any
        anon_message?: any
        archived_day?: any
        auto_archived?: any
        chat_information?: any
        home_footer?: any
        image_compress?: any
        image_save_local?: any
        login_code?: any
        password_policy?: any
        project_invite?: any
        reg?: any
        reg_identity?: any
        reg_invite?: any
        start_home?: any
        type?: any
    }

    // 邮箱设置
    export interface systemSettingEmail {
        account?: any
        ignore_addr?: any
        msg_unread_group_minute?: any
        msg_unread_user_minute?: any
        notice_msg?: any
        password?: any
        port?: any
        reg_verify?: any
        smtp_server?: any
        type?: any    

    }

}
