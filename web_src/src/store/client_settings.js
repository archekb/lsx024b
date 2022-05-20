import { defineStore } from 'pinia'


const defaultSettings = {
    "color": {
        "hight": { "value": "#D32F2FFF", "description": "Color of indicators when parameter is hight" },
        "normal": { "value": "#689F38FF", "description": "Color of indicators when parameter is normal" },
        "low": { "value": "#1976D2FF", "description": "Color of indicators when parameter is low" },
        "default": { "value": "#876AE1FF", "description": "Color of indicators which have no params" },
    },  
    "voltage": {
        "normal": {"from": 12, "to": 15, "description": "The indicator will have a normal color as long as the value is in the range", "unit": "V"},
    },
    "current": {
        "normal": {"from": 0, "to": 7, "description": "The indicator will have a normal color as long as the value is in the range", "unit": "A"},
    },
    "power": {
        "normal": {"from": 0, "to": 110, "description": "The indicator will have a normal color as long as the value is in the range", "unit": "W"},
    },
    "temperature": {
        "normal": {"from": 20, "to": 40,  "description": "The indicator will have a normal color as long as the value is in the range", "unit": "℃"},
    },
}

export const useClientSettingsStore = defineStore('clientSettings', {
    state: () => {
        return {
            _currentSettings: JSON.parse(localStorage.getItem('lsx024b_client_settings') || "{}"),
        }
    },
    actions: {
        Set(key, value) {
            var resObj = key.split(".").reduceRight((pv, cv) => ({ [cv]: pv }), value);
            this._currentSettings = merge(resObj, this._currentSettings)
            localStorage.setItem('lsx024b_client_settings', JSON.stringify(this._currentSettings));
        },
        Reset() {
            this._currentSettings = {};
            localStorage.removeItem('lsx024b_client_settings');
        }
    },
    getters: {
        settings() {
            return merge(this._currentSettings, merge(defaultSettings, {}))
        }
    }
})

// source are mutates
function merge(source, target) {
    for (const [key, val] of Object.entries(source)) {
        if (val !== null && typeof val === `object`) {
            if (target[key] === undefined) {
                target[key] = new val.__proto__.constructor();
            }
            merge(val, target[key]);
        } else {
            target[key] = val;
        }
    }
    return target; // we're replacing in-situ, so this is more for chaining than anything else
}