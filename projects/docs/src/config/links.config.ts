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
        path: 'https://github.com/minnek-digital-studio/cominnek#usage',
        label: 'Docs',
        _blank: true,
    },
};

export default linksConfig;