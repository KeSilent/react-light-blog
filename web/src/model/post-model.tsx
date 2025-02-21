
// 文章模型
export interface PostModel {
  id: string;
  title: string;
  slug: string;
  cover: string;
  description: string;
  content: string;
  state: string;
  author: {
    id: string;
    name: string;
    email: string;
    avatar: string;
  };
  category: {
    id: string;
    name: string;
  },
  keys: string[];
}