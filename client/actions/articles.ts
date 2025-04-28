"use server";

import { get } from "@/lib/api";
import {
  ArticleResponse,
  ArticlesResponse,
  FetchArticleResult,
  FetchArticlesResult,
} from "@/types/articles";

export async function getArticles(): Promise<FetchArticlesResult> {
  try {
    const response = await get<ArticlesResponse>("");

    if (response.status >= 200 && response.status < 300) {
      return { data: response.data, error: null };
    } else {
      console.error(
        "Failed to fetch articles with status:",
        response.status,
        response.data
      );
      return {
        data: null,
        error: `Failed to fetch articles (HTTP ${response.status}): ${
          response.data?.message ||
          "An unexpected error occurred from the server."
        }`,
      };
    }
  } catch (error: any) {
    console.error("Failed to fetch articles:", error);
    return {
      data: null,
      error: `An unexpected error occurred while fetching articles: ${
        error.message || "Please try again later."
      }`,
    };
  }
}

export async function getArticleById(id: string): Promise<FetchArticleResult> {
  try {
    const response = await get<ArticleResponse>(`/articles/${id}`);

    if (response.status >= 200 && response.status < 300) {
      return { data: response.data, error: null };
    } else {
      console.error(
        "Failed to fetch articles with status:",
        response.status,
        response.data
      );
      return {
        data: null,
        error: `Failed to fetch articles (HTTP ${response.status}): ${
          response.data?.message ||
          "An unexpected error occurred from the server."
        }`,
      };
    }
  } catch (error: any) {
    console.error("Failed to fetch articles:", error);
    return {
      data: null,
      error: `An unexpected error occurred while fetching articles: ${
        error.message || "Please try again later."
      }`,
    };
  }
}
