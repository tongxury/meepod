class Navigation {
    private navigation: any;

    setNavigation(navigation) {
        this.navigation = navigation
    }

    navigate(name: string, params?: { [key: string]: any }) {
        this.navigation?.navigate(name, params)
    }
}

const nav = new Navigation();

export default nav
