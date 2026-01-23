import {
    FlatList,
    Modal as RNModel,
    Pressable,
    View,
    Image,
    ActivityIndicator,
    StyleProp,
    ViewStyle
} from "react-native";
import {Card, Text, useTheme} from "react-native-paper";
import React, {useState} from "react";
import {ImageViewer as IV} from "react-native-image-zoom-viewer";
import {Avatar} from '@rneui/themed';


const ImageViewer = ({images, size, style}: {
    images: { url: string }[],
    size?: ('small' | 'medium' | 'large' | 'xlarge') | number,
    style?: StyleProp<ViewStyle>
}) => {

    const [open, setOpen] = useState<boolean>(false)

    const {colors} = useTheme()

    return <View style={style}>
        <View style={{flexDirection: "row", alignItems: "center", flexWrap: "wrap"}}>
            {
                images.map((t, i) => <View key={i} style={{margin: 2}}>
                    <Avatar
                        onPress={() => t.url && setOpen(true)}
                        size={size ?? 'large'}
                        avatarStyle={{borderRadius: 5}}
                        placeholderStyle={{backgroundColor: colors.primaryContainer}}
                        source={{uri: t.url}}
                    />
                </View>)
            }
        </View>
        <RNModel visible={open} transparent={true}>
            <IV onClick={() => setOpen(false)} imageUrls={images?.map(t => {
                return {originUrl: t.url, url: t.url}
            })}/>
        </RNModel>
    </View>

}


export default ImageViewer