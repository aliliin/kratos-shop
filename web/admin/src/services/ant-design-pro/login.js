// @ts-ignore

/* eslint-disable */
import { request } from 'umi';
/** 发送验证码 POST /api/login/captcha */

export async function getFakeCaptcha(params, options) {
  return request('/api/users/captcha', {
    method: 'GET',
    params: { ...params },
    ...(options || {}),
  });
}
