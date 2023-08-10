import { GlobalThemeOverrides } from "naive-ui"

const themeOverrides = {
    lightThemeOverrides:<GlobalThemeOverrides> {
        common: {
            primaryColor: "#84C56A",
            primaryColorHover: "#a2d98d",
            heightMedium: "32px",
            inputColor:"#F4F5F7",
            borderColor:"#F4F5F7",
        },
        Button:{
            paddingMedium:'11px 36px',
            heightMedium: "36px",
        },
        Input: {
            heightSmall: "30px",
        },
        Tag:{
            borderRadius: "4px",
            color:"rgba(167, 171, 181, 0.1)",
            textColor:"#A7ABB5",
            fontSizeMedium:"12px",
            heightMedium:'20px',
        },
        Card:{
            paddingHuge:"24px 24px 24px",
            borderRadius:"12px"
        },
        Form: {
            feedbackHeightMedium: "20px",
            feedbackFontSizeMedium: "12px"
        },
        Dialog:{
            padding:"24px",
            borderRadius:"16px",
        },
    },
    darkThemeOverrides:<GlobalThemeOverrides>  {
        common: {
            primaryColor: "#67a050",
            primaryColorHover: "#568342",
            baseColor: "#ffffff",
            heightMedium: "32px",
            inputColor:"#F4F5F7",
            borderColor:"#F4F5F7",
        },
        Input: {
            heightSmall: "30px",
        },
        Tag:{
            borderRadius: "4px",
            color:"rgba(139, 207, 112, 0.1);",
            textColor:"#8BCF70",
            fontSizeMedium:"12px",
            heightMedium:'20px',
        },
        Card:{
            paddingHuge:"24px 24px 24px",
            borderRadius:"12px"
        },
    },
}
export default themeOverrides
