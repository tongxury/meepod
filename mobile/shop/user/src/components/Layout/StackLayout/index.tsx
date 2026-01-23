import {ColorValue, FlexStyle, StyleSheet, View, ViewStyle} from "react-native";
import React from "react";

export interface StackProps extends FlexProps, BoxProps {
    children: React.ReactNode,
}

export interface BoxProps {
    spacing?: number,
    p?: number,
    m?: number,
    bg?: ColorValue | undefined;
}

export interface FlexProps {
    direction?: FlexStyle['flexDirection'];
    justify?: 'start' | 'end' | 'center' | 'between' | 'around' | 'evenly';
    items?: 'start' | 'end' | 'center' | 'stretch' | 'baseline';
    fill?: boolean | number,
    center?: boolean
}

const getValidChildren = (children: React.ReactNode) => {
    return React.Children.toArray(children).filter((child) =>
        React.isValidElement(child)
    ) as React.ReactElement[];
}

const StackLayout = ({children, direction, justify, items, fill, center, spacing, p, m, bg}: StackProps) => {

    const justifyValues = {
        'start': 'flex-start',
        'end': 'flex-end',
        'center': 'center',
        'between': 'space-between',
        'around': 'space-around',
        'evenly': 'space-evenly'
    }

    const alignValues = {
        'start': 'flex-start',
        'end': 'flex-end',
        'center': 'center',
        'stretch': 'stretch',
        'baseline': 'baseline'
    }


    const validChildren = getValidChildren(children)

    // @ts-ignore
    return <View style={{
        justifyContent: center ? 'center' : justifyValues[justify],
        alignItems: center ? 'center' : alignValues[items],
        flexDirection: direction,
        flex: typeof fill === 'boolean' ? (fill ? 1 : undefined) : fill,
        // flex: 1,
        padding: p,
        backgroundColor: bg,
        margin: m
    }}>
        {validChildren?.map((child, index) => {

                const spacingStyle = () => {
                    switch (direction) {
                        case 'column':
                            return {marginBottom: spacing};
                        case 'row':
                            return {marginEnd: spacing};
                        case 'column-reverse':
                            return {marginTop: spacing};
                        case 'row-reverse':
                            return {marginStart: spacing};
                    }
                }

                const key = typeof child.key !== 'undefined' ? child.key : index;
                const isLast = index + 1 === validChildren.length;


                return <View key={key} style={{
                    ...(child.props.style ?? {}),
                    ...(isLast ? {} : spacingStyle()),
                    backgroundColor: 'orange'
                    // marginBottom: index !== validChildren?.length - 1 ? spacing : 0,
                }}>
                    {child}
                </View>
            }
        )}
    </View>
}


export const HStack = (props: StackProps) => {
    return <StackLayout {...props} direction={'row'}/>
}

export const Stack = (props: StackProps) => {
    return <StackLayout {...props} direction={'column'}/>
}
