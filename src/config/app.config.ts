export class AppConfig {
    public static page: number = 1;
    public static ppp: number = 20;
    public static minDistance: number = 0.05;
    public static maxDistance: number = 5;
    public static timezone: any = {
        offset: '+07:00',
        tz: 'Asia/Ho_Chi_Minh'
    };
}

export enum ServerMode {
    fork = 'fork',
    cluster = 'cluster'
}
