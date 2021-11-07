import userService from '@/service/user'

const check = {
    passwordReg: /(?!^(\d+|[a-zA-Z]+|[~!@#$%^&*?]+)$)^[\w~!@#$%^&*?]{6,12}$/,
    mobileReg: /^\d{11}$/,
    emailReg: /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/,
    validateName: (parmas) => {
        userService.validateName(parmas)
    }
}

export default check