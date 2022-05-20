import axios from 'axios';
import { defineStore } from 'pinia'

export const useControllerStateStore = defineStore('controllerState', {
    state: () => {
        return {
            _timer: null,
            state: {},
            inProgress: false,
            error: ""
        }
    },
    actions: {
        start() {
            this.inProgress = true;
            axios
                .create({
                    baseURL: process.env.NODE_ENV == "development" ? process.env.VUE_APP_BACK : "",
                    timeout: 5000,
                    headers: {
                        'Content-Type': 'application/json',
                    }
                })
                .get("/api/state")
                .then(res => {
                    this.inProgress = false;
                    if (res.status != 200 ) throw "Wrong Server answer";
                    if (res.data) {
                        this.state = res.data;
                    } else {
                        throw "Unknown answer";
                    }
                })
                .catch(e => {
                    this.inProgress = false;
                    if (e.request) {
                        this.error = e.message;
                    } else {
                        this.error = e;
                    }
                })
            this._timer = setTimeout(this.start, this.state?.update_interval * 1000 || 3000);
        },
        stop() {
            clearTimeout(this._timer)
        }
    }
})