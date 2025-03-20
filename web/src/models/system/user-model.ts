//用户模型
export interface UserModel {
  id?: number;
  username?: string;
  nickName?: string;
  avatar?: string;
  phone?: string;
  email?: string;
  status?: boolean;
  headerImg?: string;
}

export interface ChangePasswordReq {
  password?: string;
  newPassword?: string;
}