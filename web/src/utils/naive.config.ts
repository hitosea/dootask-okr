import { datePickerDark } from "naive-ui"

const themeOverrides = {
    lightThemeOverrides: {
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
    },
    darkThemeOverrides: {
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
    },
}
export default themeOverrides
