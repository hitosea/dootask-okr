import utils from "./utils";
import * as localforage from "localforage";

localforage.config({name: 'web', storeName: 'common'});

const local = {
    __timer: {},

    save(key, value, delay = 100) {
        if (typeof this.__timer[key] !== "undefined") {
            clearTimeout(this.__timer[key])
            delete this.__timer[key]
        }
        this.__timer[key] = setTimeout(async _ => {
            await localforage.setItem(key, value)
        }, delay)
    },

    del(key) {
        localforage.removeItem(key).then(_ => {
        })
    },

    set(key, value) {
        return localforage.setItem(key, value)
    },

    remove(key) {
        return localforage.removeItem(key)
    },

    clear() {
        return localforage.clear()
    },

    value(key) {
        return localforage.getItem(key)
    },

    async string(key, def = "") {
        const value = await this.value(key)
        return typeof value === "string" || typeof value === "number" ? value : def;
    },

    async int(key, def = 0) {
        const value = await this.value(key)
        return typeof value === "number" ? value : def;
    },

    async boolean(key, def = false) {
        const value = await this.value(key)
        return typeof value === "boolean" ? value : def;
    },

    async array(key, def = []) {
        const value = await this.value(key)
        return utils.isArray(value) ? value : def;
    },

    async json(key, def = {}) {
        const value = await this.value(key)
        return utils.isJson(value) ? value : def;
    },
}

export default local
