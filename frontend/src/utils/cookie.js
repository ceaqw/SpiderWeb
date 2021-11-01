// cookie操作

import Cookies from 'js-cookie'

export function getByKey(key) {
    return Cookies.get(key)
}

export function setByKey(key, value) {
    return Cookies.set(key, value)
}

export function removeByKey(key) {
    return Cookies.remove(key)
}