This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```
  

# 博客系统开发文档

**版本**：v1.0

**日期**：2025-01-07

**作者**：用户

  

---

  

## 一、项目简介

**项目名称**：自建博客系统 + Obsidian自动发布插件

**目标**：开发一个简单、高效的博客系统，逐步替代现有的WordPress，实现自动发布和SEO优化。博客系统将对接AI功能，辅助自动生成摘要、关键词和文章分类标签，提高运营效率。

  

**技术栈**：

- **前端**：Next.js（React框架）
  - [echarts](https://echarts.apache.org/zh/index.html)
  - [Markdown编辑器](https://github.com/uiwjs/react-md-editor)
  - [HTML防止XSS攻击](https://github.com/rehypejs/rehype-sanitize)
  - [Markdown处理官方文档](https://nextjs.org/docs/pages/building-your-application/configuring/mdx#install-dependencies)
  - [ejs模板引擎](https://ejs.bootcss.com/)

- **后端**：Next.js API路由 + Node.js

- **文章管理**：Markdown + Obsidian插件

- **数据库**：SQLite / PostgreSQL（选配）

- **AI对接**：OpenAI / Claude / 其他可配置的AI API

- **部署方式**：Vercel / Netlify / VPS

  

---

  

## 二、核心功能设计

  

### 1. 博客系统功能模块

| 模块 | 功能说明 | 备注 |
| ---------------------- | ------------------------------- | ---------------------------- |
| **文章管理** | 支持Markdown格式写作，自动解析frontmatter，静态生成HTML页面 | 基于Next.js静态生成（SSG） |
| **文章发布** | 通过Obsidian插件一键发布Markdown文章至博客系统 | 自动生成静态页面 |
| **SEO优化** | 自动生成meta标签（title、description、keywords），优化爬虫抓取 | 动态生成或AI辅助生成 |
| **首页文章列表** | 展示最新文章列表，分页展示，支持搜索和分类 | 可选动态或静态生成 |
| **文章详情页** | 访问独立文章页面，URL格式为 `/articles/[slug]` | 静态生成 |
| **标签和分类** | 支持文章标签与分类系统，便于读者筛选和归档 | |
| **站点地图（sitemap）** | 自动生成`sitemap.xml`，方便爬虫抓取全站文章 | 自动定期更新 |
| **RSS订阅** | 自动生成RSS订阅源，方便用户订阅新文章 | |
| **评论系统** | 集成第三方评论系统（如Disqus或自建评论系统） | 可选配置 |

---

### 2. Obsidian插件功能

| 功能 | 说明 |
|----------------------|----------------------------------------------------------------|
| **文章同步** | 将Obsidian中的Markdown文章一键同步到博客系统 |
| **自动发布** | 发布后自动触发Next.js重新构建静态页面，实现内容自动更新 |
| **AI摘要生成** | 文章发布时自动调用AI生成文章摘要，添加至frontmatter |
| **关键词提取** | 自动提取文章关键词，更新meta标签，提高SEO效果 |

---

## 三、系统架构设计

### 1. 架构概览

```
+----------------------------+
|        用户浏览器           |
+----------------------------+
           ↓                  
+----------------------------+
|        Next.js前端          |  
| - 文章详情页面（SSG）       |  
| - 首页文章列表（SSG）       |  
+----------------------------+
           ↓                  
+----------------------------+
|       Next.js API路由        |  
| - Obsidian文章接收接口       |  
| - AI服务调用接口            |  
+----------------------------+
           ↓                  
+----------------------------+
|       数据存储层（选配）     |  
| - Markdown文件              |  
| - SQLite/PostgreSQL         |  
+----------------------------+
           ↓                  
+----------------------------+
|        部署平台             |  
| - Vercel/Netlify/VPS        |  
+----------------------------+

```

---

## 四、静态页面生成策略

### 静态生成方式（SSG）

- 使用Next.js的`getStaticProps`和`getStaticPaths`生成文章静态页面。

- 文章列表和详情页均在构建时生成HTML，保证SEO友好。

- 构建流程在每次文章发布或修改后自动触发。
  

**代码示例**：

```tsx
// pages/articles/[slug].tsx

import { GetStaticPaths, GetStaticProps } from 'next';

import fs from 'fs';

import path from 'path';

import matter from 'gray-matter';

export default function ArticlePage({ frontmatter, content }) {
  return (
    <div>
      <h1>{frontmatter.title}</h1>

      <p>{frontmatter.date}</p>

      <article dangerouslySetInnerHTML={{ __html: content }}></article>
    </div>
  );
}

// 生成静态路径

export const getStaticPaths: GetStaticPaths = async () => {
  const files = fs.readdirSync(path.join('articles'));

  const paths = files.map((filename) => ({
    params: { slug: filename.replace('.md', '') },
  }));

  return { paths, fallback: false };
};

// 读取Markdown文件内容

export const getStaticProps: GetStaticProps = async ({ params }) => {
  const filePath = path.join('articles', `${params?.slug}.md`);

  const fileContent = fs.readFileSync(filePath, 'utf-8');

  const { data: frontmatter, content } = matter(fileContent);

  return {
    props: {
      frontmatter,

      content,
    },
  };
};
```

## 五、AI集成功能
### 1. AI自动摘要生成

- 调用AI接口，为文章生成简要摘要，自动填充到文章frontmatter中。

**示例代码**：
```tsx
const summary = await openai.createCompletion({
  prompt: `生成以下文章的摘要:\n\n${content}`,
  model: 'gpt-4',
  max_tokens: 150,
});
```

### 2. AI关键词提取

- 自动提取文章核心关键词，填充meta标签，提高文章在搜索引擎的可见度。
```tsx
const keywords = await openai.createCompletion({
  prompt: `提取以下文章的关键词:\n\n${content}`,
  model: 'gpt-4',
  max_tokens: 100,
});
```

## 六、SEO优化策略

### 1. 自动生成Meta标签

- 文章标题、描述和关键词自动填充到meta标签中，提高爬虫抓取效率。
```tsx
import Head from 'next/head';

export default function Article({ frontmatter }) {
  return (
    <>
      <Head>
        <title>{frontmatter.title}</title>
        <meta name="description" content={frontmatter.excerpt} />
        <meta name="keywords" content={frontmatter.keywords} />
      </Head>
      <h1>{frontmatter.title}</h1>
    </>
  );
}
```

### 2. 自动生成sitemap.xml
```tsx
const files = fs.readdirSync(path.join('articles'));
const urls = files.map((filename) => {
  return `<url><loc>https://yourblog.com/articles/${filename.replace('.md', '')}</loc></url>`;
}).join('');

const sitemap = `<?xml version="1.0"?><urlset>${urls}</urlset>`;
fs.writeFileSync('public/sitemap.xml', sitemap);
```

## 七、部署与上线

- **托管平台**：Vercel（Next.js原生支持）或Netlify
- **自动构建流程**：每次Markdown文件更新后触发自动构建，生成新的静态页面。
- **CDN缓存**：静态页面托管在CDN上，全球加速，提升访问速度。

## 八、后续规划

- 增加评论系统
- 支持文章置顶与推荐
- 增加Obsidian中的草稿管理功能
- 集成多语言支持