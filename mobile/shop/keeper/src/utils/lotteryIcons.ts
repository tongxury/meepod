export const lotteryIcons = {
    'ssq': require('../assets/ssq.png'),
    'f3d': require('../assets/f3d.png'),
    'x7c': require('../assets/x7c.png'),
    'rx9': require('../assets/rx9.png'),
    'sfc': require('../assets/sfc.png'),
    'zjc': require('../assets/zjc.png'),
    'pl3': require('../assets/pl3.png'),
    'pl5': require('../assets/pl5.png'),
    'kl8': require('../assets/kl8.png'),
    'dlt': require('../assets/ssq.png'),
};

export const getLotteryIcon = (itemId: string, fallbackUrl?: string) => {
    return lotteryIcons[itemId] || (fallbackUrl ? { uri: fallbackUrl } : null);
};
