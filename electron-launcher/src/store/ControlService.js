import * as EventEmitter from 'events';

let instance = null;
class ControlService extends EventEmitter {

    constructor(electron) {
        super();                        
        this.promiseIpc = window.promiseIpc;
    }

    init(store) {                
        this.promiseIpc.send('ready', 'OK');        
        instance.$store = store;                                  
    }    

    async connect(args) {
        return this.promiseIpc.send('connect', {
            host: args.host,
            user: args.user,
            password: args.password,
            sshKeyAuth: args.sshKeyAuth,
            keyfileLocation: args.keyfileLocation,
            stereumRelease: args.stereumRelease,
        });
    }

    async disconnect(args) {
        return this.promiseIpc.send('disconnect', args);
    }

    async inquire(args) {
        return this.promiseIpc.send('inquire', args);
    }

    async setup(args) {
        return this.promiseIpc.send('setup', args);
    }

    async openTunnels(args) {
        await this.promiseIpc.send('tunnel', args);
        return args;
    }

}
if (!instance) {
    instance = new ControlService(window.electron);
}
export default instance;
