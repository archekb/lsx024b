export const nameFormatter = (str) => {
    return str.split('_').map(x => x[0].toUpperCase() + x.slice(1)).join(' ')
}

export const unitFormatter = (name, value) => {
    name = name.toLowerCase();
    if (typeof value === 'number' && isFinite(value)) {
        if (name.includes("temperature")) return "℃";
        if (name.includes("voltage")) return "V";
        if (name.includes("power")) return "W";
        if (name.includes("current")) return "A";
        if (name.includes("energy")) return "KWh"
        if (name.includes("soc")) return "%"
    }
    return ""
}

export const colorFormatter = (settings, name, value, defaultValue) => {
    defaultValue = defaultValue || settings?.color?.default?.value;
    name = name.toLowerCase();
 
    if (typeof value === 'number' && isFinite(value)) {
        if (name.includes("temperature")) {
            if (settings?.temperature?.normal?.from <= value && value <= settings?.temperature?.normal?.to) {
                return settings?.color?.normal?.value || defaultValue;
            } else if (value > settings?.temperature?.normal?.to) {
                return settings?.color?.hight?.value || defaultValue;
            } else if (value < settings?.temperature?.normal?.from) {
                return settings?.color?.low?.value || defaultValue;
            }
        }
        
        if (name.includes("voltage")) {
            if (settings?.voltage?.normal?.from <= value && value <= settings?.voltage?.normal?.to) {
                return settings?.color?.normal?.value || defaultValue;
            } else if (value > settings?.voltage?.normal?.to) {
                return settings?.color?.hight?.value || defaultValue;
            } else if (value < settings?.voltage?.normal?.from) {
                return settings?.color?.low?.value || defaultValue;
            }
        }

        if (name.includes("power")) {
            if (settings?.power?.normal?.from <= value && value <= settings?.power?.normal?.to) {
                return settings?.color?.normal?.value || defaultValue;
            } else if (value > settings?.power?.normal?.to) {
                return settings?.color?.hight?.value || defaultValue;
            } else if (value < settings?.power?.normal?.from) {
                return settings?.color?.low?.value || defaultValue;
            }
        }

        if (name.includes("current")) {
            if (settings?.current?.normal?.from <= value && value <= settings?.current?.normal?.to) {
                return settings?.color?.normal?.value || defaultValue;
            } else if (value > settings?.current?.normal?.to) {
                return settings?.color?.hight?.value || defaultValue;
            } else if (value < settings?.current?.normal?.from) {
                return settings?.color?.low?.value || defaultValue;
            }
        }
    }

    return defaultValue;
}
