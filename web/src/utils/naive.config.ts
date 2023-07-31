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
        Input: {
            heightSmall: "30px",
        },
        Tag:{
            borderRadius: "4px",
            color:"rgba(167, 171, 181, 0.1)",
            textColor:"#A7ABB5",
        },
        Card:{
            paddingHuge:"24px 24px 24px",
            borderRadius:"12px"
        },
        Form: {
            feedbackHeightMedium: "20px",
            feedbackFontSizeMedium: "12px"
        }
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
            color:"rgba(167, 171, 181, 0.1)",
            textColor:"#A7ABB5",
        },
        Card:{
            paddingHuge:"24px 24px 24px",
            borderRadius:"12px"
        },
    },
}
export default themeOverrides
