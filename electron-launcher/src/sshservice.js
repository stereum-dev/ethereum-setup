const { readFileSync } = require('fs');
const { Client } = require('ssh2');
const tunnel = require('tunnel-ssh');
const fs = require('fs')

export class SSHService {

    constructor() { 
        this.conn = null;
        this.connectionInfo = null;
        this.connected = false;
    }

    async connect(connectionInfo) {        
        this.connectionInfo = connectionInfo;
        this.conn = new Client();

        console.log('CONNECT: connectionInfo', this.connectionInfo)

        return new Promise((resolve, reject) => {            
            this.conn.on('error', (error) => {
                this.conn.end();
                reject(error);
            });
            this.conn.on('ready', () => {        
                this.connected = true;
                resolve(this.conn);
            }).connect({
                host: connectionInfo.host,
                port: connectionInfo.port || 22,
                username: connectionInfo.user || 'root',
                password: connectionInfo.password || undefined,
                privateKey: connectionInfo.privateKey || undefined,
            });
        })        
    }

    async disconnect(connectionInfo) {        
        
        console.log('DISCONNECT: connectionInfo', this.connectionInfo)
        
        return new Promise((resolve, reject) => {
            try {
                this.conn.end();
                this.connected = false;
                resolve(true);
            } catch(error) {
                reject(error);
            }            
        })
    }
        
    async exec(command) {        

        console.log('exec', command)        
        
        return new Promise((resolve, reject) => {
            let data = {
                rc: -1,
                stdout: '',
                stderr: '',
            }                        
            this.conn.exec(command, (err, stream) => {
                if (err) return reject(err);
                stream
                    .on('close', (code, signal) => {                                                
                        console.log('stream closed', code);                             
                        data.rc = code;
                        resolve(data);
                    })                        
                    .on('data', (stdout) => { 
                        console.log('stdout got data', stdout.toString('utf8'));                             
                        data.stdout = stdout.toString('utf8');
                    })
                    .stderr.on('data', (stderr) => {
                        console.log('stderr got data', stderr.toString('utf8'));                             
                        data.stderr = stderr.toString('utf8');
                    });                
            })
        });
    }

    async tunnel(tunnelConfig) {
        return new Promise((resolve, reject) => {
            var config = {
                keepAlive: true,
                username: this.connectionInfo.user || 'root',
                password: this.connectionInfo.password,
                host: this.connectionInfo.host,
                port: this.connectionInfo.port || 22,
                dstHost: 'localhost',
                dstPort: tunnelConfig.dstPort,
                localHost:'0.0.0.0',
                localPort: tunnelConfig.localPort,
                privateKey: this.connectionInfo.privateKey || undefined,                
            };
            console.log(config);
            
            var server = tunnel(config, (error, server) => {
                if (error) {
                    console.error('Tunnel Connection failed!');
                    return reject(error);
                }
                console.error(`Tunnel Connection established! (${tunnelConfig.localPort})`);
                resolve(server);
            });

            server.on('error', function(error) {
                console.error('Tunnel connection error: ', error);
            });
        });
    }

}