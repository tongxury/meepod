import {Provider as PaperProvider} from 'react-native-paper';
import {Provider as AntProvider} from '@ant-design/react-native'
import {PartialTheme} from "@ant-design/react-native/lib/style";
import {MD3LightTheme as DefaultTheme, MD3DarkTheme} from 'react-native-paper';
import {ThemeProvider as RneThemeProvider, Button, createTheme} from '@rneui/themed';
import {useState} from "react";

const lightColors = {
    "primary": "rgb(245,91,3)",
    "onPrimary": "rgb(255, 255, 255)",
    "primaryContainer": "rgb(248,218,201)",
    "onPrimaryContainer": "rgb(245,91,3)",

    "secondary": "rgb(102, 90, 111)",
    "onSecondary": "rgb(255, 255, 255)",
    "secondaryContainer": "rgb(236,231,231)",
    "onSecondaryContainer": "rgb(64,68,68)",

    "tertiary": "rgb(128, 81, 88)",
    "onTertiary": "rgb(255, 255, 255)",
    "tertiaryContainer": "rgb(255, 217, 221)",
    "onTertiaryContainer": "rgb(50, 16, 23)",

    "error": "rgb(186, 26, 26)",
    "onError": "rgb(255, 255, 255)",
    "errorContainer": "rgb(255, 218, 214)",
    "onErrorContainer": "rgb(65, 0, 2)",

    "background": "rgb(253,253,253)",
    "onBackground": "rgb(29, 27, 30)",

    "surface": "rgb(255, 251, 255)",
    "onSurface": "rgb(29, 27, 30)",
    "surfaceVariant": "rgb(229,227,226)",
    "onSurfaceVariant": "rgb(74, 69, 78)",
    "outline": "rgb(248,218,201)",
    "outlineVariant": "rgb(248,218,201)",
    "shadow": "rgb(0, 0, 0)",
    "scrim": "rgb(0, 0, 0)",
    "inverseSurface": "rgb(50, 47, 51)",
    "inverseOnSurface": "rgb(248,218,201)",
    "inversePrimary": "rgb(248,218,201)",
    "elevation": {
        "level0": "transparent",
        "level1": "rgb(255,255,255)",
        "level2": "rgb(250,228,217)", // 底部菜单背景
        "level3": "rgb(246,228,218)",
        "level4": "rgb(246,229,219)",
        "level5": "rgb(246,230,221)"
    },
    "surfaceDisabled": "rgba(29, 27, 30, 0.12)",
    "onSurfaceDisabled": "rgba(29, 27, 30, 0.38)",
    "backdrop": "rgba(51, 47, 55, 0.4)"
}
const darkColors = {
    "primary": "rgb(220, 184, 255)",
    "onPrimary": "rgb(71, 12, 122)",
    "primaryContainer": "rgb(95, 43, 146)",
    "onPrimaryContainer": "rgb(240, 219, 255)",

    "secondary": "rgb(208, 193, 218)",
    "onSecondary": "rgb(54, 44, 63)",
    "secondaryContainer": "rgb(77, 67, 87)",
    "onSecondaryContainer": "rgb(237, 221, 246)",

    "tertiary": "rgb(243, 183, 190)",
    "onTertiary": "rgb(75, 37, 43)",
    "tertiaryContainer": "rgb(101, 58, 65)",
    "onTertiaryContainer": "rgb(255, 217, 221)",

    "error": "rgb(255, 180, 171)",
    "onError": "rgb(105, 0, 5)",
    "errorContainer": "rgb(147, 0, 10)",
    "onErrorContainer": "rgb(255, 180, 171)",

    "background": "rgb(29, 27, 30)",
    "onBackground": "rgb(231, 225, 229)",

    "surface": "rgb(29, 27, 30)",
    "onSurface": "rgb(231, 225, 229)",
    "surfaceVariant": "rgb(74, 69, 78)",
    "onSurfaceVariant": "rgb(204, 196, 206)",

    "outline": "rgb(150, 142, 152)",
    "outlineVariant": "rgb(74, 69, 78)",
    "shadow": "rgb(0, 0, 0)",
    "scrim": "rgb(0, 0, 0)",
    "inverseSurface": "rgb(231, 225, 229)",
    "inverseOnSurface": "rgb(50, 47, 51)",
    "inversePrimary": "rgb(120, 69, 172)",
    "elevation": {
        "level0": "transparent",
        "level1": "rgb(39, 35, 41)",
        "level2": "rgb(44, 40, 48)",
        "level3": "rgb(50, 44, 55)",
        "level4": "rgb(52, 46, 57)",
        "level5": "rgb(56, 49, 62)"
    },
    "surfaceDisabled": "rgba(231, 225, 229, 0.12)",
    "onSurfaceDisabled": "rgba(231, 225, 229, 0.38)",
    "backdrop": "rgba(51, 47, 55, 0.4)"
}
const ThemeProvider = ({children}) => {


    // const colorScheme = useColorScheme();
    const [theme, setTheme] = useState<'light' | 'dark'>('light')

    const currentColors =
        theme === 'dark' ? darkColors : lightColors

    const paperTheme = {
        ...DefaultTheme,
        roundness: 1,
        colors: {
            ...DefaultTheme.colors,
            ...currentColors
        },
    };

    const antdTheme: PartialTheme = {
        // 支付宝钱包默认主题
        // https://github.com/ant-design/ant-design-mobile/wiki/设计变量表及命名规范

        // 色彩, NOTE: must use `#000000` instead of `#000`
        // https://facebook.github.io/react-native/docs/colors.html
        // 8-digit-hex to 4-digit hex https://css-tricks.com/8-digit-hex-codes/
        // https://www.chromestatus.com/feature/5685348285808640 chrome will support `#RGBA`
        // 文字色
        color_text_base: '#000000', // 基本
        color_text_base_inverse: '#ffffff', // 基本 _ 反色
        color_text_secondary: '#a4a9b0', // 辅助色
        color_text_placeholder: '#bbbbbb', // 文本框提示
        color_text_disabled: '#bbbbbb', // 失效
        color_text_caption: '#888888', // 辅助描述
        color_text_paragraph: '#333333', // 段落
        color_link: currentColors.primary, // 链接

        // 背景色
        fill_base: '#ffffff', // 组件默认背景
        fill_body: '#f5f5f9', // 页面背景
        fill_tap: '#dddddd', // 组件默认背景 _ 按下
        fill_disabled: '#dddddd', // 通用失效背景
        fill_mask: 'rgba(0, 0, 0, .4)', // 遮罩背景
        color_icon_base: '#cccccc', // 许多小图标的背景，比如一些小圆点，加减号
        fill_grey: '#f7f7f7',

        // 透明度
        opacity_disabled: '0.3', // switch checkbox radio 等组件禁用的透明度

        // 全局/品牌色
        brand_primary: currentColors.primary,
        brand_primary_tap: currentColors.secondary,
        brand_success: '#6abf47',
        brand_warning: '#f4333c',
        brand_error: '#f4333c',
        brand_important: '#ff5b05', // 用于小红点
        brand_wait: '#108ee9',

        // 边框色
        border_color_base: '#dddddd',

        // 字体尺寸
        // ---
        font_size_icontext: 10,
        font_size_caption_sm: 12,
        font_size_base: 14,
        font_size_subhead: 15,
        font_size_caption: 16,
        font_size_heading: 17,

        // 圆角
        // ---
        radius_xs: 2,
        radius_sm: 3,
        radius_md: 5,
        radius_lg: 7,

        // 边框尺寸
        // ---
        border_width_sm: 0.5,
        border_width_md: 1,
        border_width_lg: 2,

        // 间距
        // ---
        // 水平间距
        h_spacing_sm: 5,
        h_spacing_md: 8,
        h_spacing_lg: 15,

        // 垂直间距
        v_spacing_xs: 3,
        v_spacing_sm: 6,
        v_spacing_md: 9,
        v_spacing_lg: 15,
        v_spacing_xl: 21,

        // 高度
        // ---
        line_height_base: 1, // 单行行高
        line_height_paragraph: 1.5, // 多行行高

        // 图标尺寸
        // ---
        icon_size_xxs: 15,
        icon_size_xs: 18,
        icon_size_sm: 21,
        icon_size_md: 22, // 导航条上的图标
        icon_size_lg: 36,

        // 动画缓动
        // ---
        ease_in_out_quint: 'cubic_bezier(0.86, 0, 0.07, 1)',

        // 组件变量
        // ---

        actionsheet_item_height: 50,
        actionsheet_item_font_size: 18,

        // button
        button_height: 47,
        button_font_size: 18,

        button_height_sm: 23,
        button_font_size_sm: 12,

        primary_button_fill: currentColors.primary,
        primary_button_fill_tap: '#0e80d2',

        ghost_button_color: currentColors.primary, // 同时应用于背景、文字颜色、边框色
        ghost_button_fill_tap: `${currentColors.primary}99`, // alpha 60%  https://codepen.io/chriscoyier/pen/XjbzAW

        warning_button_fill: '#e94f4f',
        warning_button_fill_tap: '#d24747',

        link_button_fill_tap: '#dddddd',
        link_button_font_size: 16,

        // modal
        modal_font_size_heading: 18,
        modal_button_font_size: 18, // 按钮字号
        modal_button_height: 50, // 按钮高度

        // list
        list_title_height: 30,
        list_item_height_sm: 35,
        list_item_height: 44,

        // input
        input_label_width: 17, // InputItem、TextareaItem 文字长度基础值
        input_font_size: 17,
        input_color_icon: '#cccccc',
        input_color_icon_tap: currentColors.primary,

        // tabs
        tabs_color: currentColors.primary,
        tabs_height: 42,
        tabs_font_size_heading: 15,

        // segmented_control
        segmented_control_color: currentColors.primary, // 同时应用于背景、文字颜色、边框色
        segmented_control_height: 27,
        segmented_control_fill_tap: `${currentColors.primary}10`,

        // tab_bar
        tab_bar_fill: '#ebeeef',
        tab_bar_height: 50,

        // toast
        toast_fill: 'rgba(0, 0, 0, .8)',

        // search_bar
        search_bar_fill: '#efeff4',
        search_bar_height: 44,
        search_bar_input_height: 28,
        search_bar_font_size: 15,
        search_color_icon: '#bbbbbb', // input search icon 的背景色

        // notice_bar
        notice_bar_fill: '#fffada',
        notice_bar_height: 36,

        // checkbox
        checkbox_fill: '#1890ff',
        checkbox_fill_disabled: '#f5f5f5',
        checkbox_border: '#d9d9d9',
        checkbox_border_disabled: '#b9b9b9',

        // switch
        switch_fill: '#1890ff',
        switch_unchecked: '#cccccc',
        switch_unchecked_disabled: '#cccccc66', // switch_fill的40%透明度
        switch_checked_disabled: '#1890ff66', // switch_unchecked的40%透明度

        // tag
        tag_height: 25,
        tag_small_height: 15,

        // picker
        option_height: 42, // picker 标题的高度

        toast_zindex: 1999,
        action_sheet_zindex: 1000,
        popup_zindex: 999,
        modal_zindex: 999,
    }

    const renTheme = createTheme({
        lightColors: {
            primary: currentColors.primary,
            secondary: currentColors.secondaryContainer,
        },
        components: {
            Skeleton: {
                skeletonStyle: {backgroundColor: currentColors.surface},
                style: {backgroundColor: currentColors.surfaceVariant},
                animation: 'wave'
            },
            Button: {
                titleStyle: {fontSize: 14, marginHorizontal: 10},
                radius: 5,
            },
            Chip: {
                radius: 5,
                color: currentColors.secondaryContainer,
                buttonStyle: {padding: 1},
                titleStyle: {
                    color: currentColors.onPrimary
                },
                // size: "sm",
            },
            CheckBox: {
                titleProps: {style: {marginHorizontal: 0, alignContent: "center"}},
                containerStyle: {margin: 0, padding: 0, backgroundColor: currentColors.background}
            }
        },
    });

    return <RneThemeProvider theme={renTheme}>
        <PaperProvider theme={paperTheme}>
            <AntProvider theme={antdTheme}>
                {children}
            </AntProvider>
        </PaperProvider>
    </RneThemeProvider>
}

export default ThemeProvider