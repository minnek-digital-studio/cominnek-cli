interface ILinksConfig {
    [key: string]: {
        path: string;
        label: string;
        hidden?: boolean;
        _blank?: boolean;
    };
}

const linksConfig: ILinksConfig = {
    home: {
        path: '/',
        label: 'Home',
        hidden: true,
    },
    docs: {
        path: '/docs',
        label: 'Docs',
    },
};

export default linksConfig;