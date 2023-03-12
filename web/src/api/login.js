import { request } from '@/utils/request.js'

export default {


  /**
   * 获取验证码
   * @returns
   */
  getCaptch(params={}) {
    return request({
      url: 'base/captcha',
      method: 'post',
      data:params
    })
  },

  /**
   * 用户登录
   * @param {object} params
   * @returns
   */
  login(params = {}) {
    return request({
      url: 'base/login',
      method: 'post',
      data: params
    })
  },

  /**
   * 用户退出
   * @param {object} params
   * @returns
   */
  logout(params = {}) {
    return request({
      url: 'system/logout',
      method: 'post',
      data: params
    })
  },

  /**
   * 获取登录用户信息
   * @param {object} params
   * @returns
   */
  getInfo(params = {}) {
    return request({
      url: 'user/getUserInfo',
      method: 'get',
      data: params
    })
  }
}
