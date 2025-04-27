export type Article = {
  ID: number;
  Image: string;
  Title: string;
  Content: string;
  CreatedAt: string;
  User: {
    ID: number;
    Image: string;
    Name: string;
    Email: string;
  };
};

export type ArticlesResponse = {
  message: string;
  featured: Article[];
  stories: Article[];
  status: number;
};

export type ArticleResponse = {
  article: Article;
  message: string;
  status: number;
};

export type FetchArticlesResult =
  | { data: ArticlesResponse; error: null }
  | { data: null; error: string };

export type FetchArticleResult =
  | { data: ArticleResponse; error: null }
  | { data: null; error: string };
