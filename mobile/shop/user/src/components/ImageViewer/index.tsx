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
import {Card, Text} from "react-native-paper";
import React, {useState} from "react";
import {ImageViewer as IV} from "react-native-image-zoom-viewer";
import {Avatar} from '@rneui/themed';


const ImageViewer = ({images, size, style}: {
    images: { url: string }[],
    size?: ('small' | 'medium' | 'large' | 'xlarge') | number,
    style?: StyleProp<ViewStyle>
}) => {

    const [open, setOpen] = useState<boolean>(false)

    return <View style={style}>
        <View style={{flexDirection: "row", alignItems: "center", flexWrap: "wrap"}}>
            {
                images.map((t, i) => <View key={i} style={{margin: 2}}>
                    <Avatar
                        onPress={() => setOpen(true)}
                        size={size ?? 'large'}
                        avatarStyle={{borderRadius: 5}}
                        source={{uri: t.url}}
                    />
                </View>)
            }
        </View>
        {/*<FlatList*/}
        {/*    data={images}*/}
        {/*    numColumns={4}*/}
        {/*    renderItem={({item: x}) => {*/}
        {/*        // return <Card style={{flex: 1}} onPress={() => setOpen(true)}>*/}
        {/*        //     <Card.Cover source={{uri: x.url}}/>*/}
        {/*        // </Card>*/}
        {/*        return <View style={{flex: 1}}>*/}
        {/*            <Avatar*/}
        {/*                avatarStyle={{borderRadius: 5}}*/}
        {/*                onPress={() => setOpen(true)}*/}
        {/*                size="large"*/}
        {/*                source={{uri: x.url}}*/}
        {/*            />*/}
        {/*        </View>*/}
        {/*    }}*/}
        {/*    keyExtractor={(x) => x.url}*/}
        {/*/>*/}
        <RNModel visible={open} transparent={true}>
            <IV onClick={() => setOpen(false)} imageUrls={images?.map(t => {
                return {originUrl: t.url, url: t.url}
            })}/>
        </RNModel>
    </View>

}


export default ImageViewer