import { GlobalThemeOverrides } from "naive-ui"

const themeOverrides = {
    lightThemeOverrides:<GlobalThemeOverrides> {
        common: {
            primaryColor: "#84C56A",
            primaryColorHover: "#a2d98d",
            heightMedium: "32px",
            heightLarge: "36px",
            borderColor:"#e8e8e8",
            borderRadius:"4px",
            textColor2:"#606266",
            errorColor:"#FF0000",
        },
        Button:{
            paddingMedium:'11px 36px',
            paddingSmall:'0 15px',
            heightSmall: "32px",
            heightMedium: "36px",
            fontWeightStrong: "400",
        },
        Input: {
            heightSmall: "30px",
            borderRadius: "4px",
            color:"#fff",
            colorFocus:"#fff",
            border:"1px solid #e8e8e8",
        },
        Tag:{
            borderRadius: "4px",
            color:"rgba(167, 171, 181, 0.1)",
            textColor:"#A7ABB5",
            fontSizeMedium:"12px",
            heightMedium:'24px',
            textColorError:"#F5222D",
            colorError:"#FFF1FO",
        },
        Card:{
            paddingHuge:"24px 24px 24px",
            borderRadius:"12px",
            titleTextColor:"#303133",
        },
        Form: {
            feedbackHeightMedium: "16px",
            feedbackFontSizeMedium: "12px",
            labelTextColor:"#606266",
            feedbackTextColorError:"#FF0000",
            asteriskColor:"#FF0000",
        },
        Dialog:{
            padding:"24px",
            borderRadius:"16px",
        },
        Radio:{
            dotColorDisabled:'#C0C0C0',
            textColor:"#606266",
        },
        Checkbox:{
            borderRadius:'99px',
            sizeMedium:'18px',
            border:'1px solid #B1B3B7',
        },
        DataTable:{
            thColor:"#F8F8F9",
            thTextColor:"#606266",
            thPaddingMedium:"8px",
            tdPaddingMedium:"8px",
            thPaddingLarge:"8px",
            emptyPadding:'13px 0'
        },
        Popover:{
            borderRadius:"8px",
        },
        Tooltip:{
            color:"rgba(81, 90, 110, 0.75)",
        },
        Tabs:{
            tabFontWeightActive:"500",
            tabTextColorLine:"#606266",
        },
        Drawer:{
            bodyPadding:"16px 24px 24px 24px",
        },
        Pagination:{
            buttonColor:"#ffffff"
        },
    },
    darkThemeOverrides:<GlobalThemeOverrides>  {
        common: {
            primaryColor: "#8bcf70",
            primaryColorHover: "#a2d98d",
        },
        Button:{
            textColorPrimary:'#000',
            textColorHoverPrimary:'#000',
            textColorPressedPrimary:'#000',
            textColorFocusPrimary:'#000',
            textColorDisabledPrimary:'#000',
        },
        // common: {
        //     primaryColor: "#67a050",
        //     primaryColorHover: "#568342",
        //     baseColor: "#ffffff",
        //     heightMedium: "32px",
        //     inputColor:"#F4F5F7",
        //     borderColor:"#F4F5F7",
        // },
        // Input: {
        //     heightSmall: "30px",
        // },
        // Tag:{
        //     borderRadius: "4px",
        //     color:"rgba(139, 207, 112, 0.1);",
        //     textColor:"#8BCF70",
        //     fontSizeMedium:"12px",
        //     heightMedium:'20px',
        // },
        // Card:{
        //     paddingHuge:"24px 24px 24px",
        //     borderRadius:"12px"
        // },
    },
}
export default themeOverrides
