"use server";

import { get } from "@/lib/api";

type Article = {
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

type ArticlesResponse = {
  featured: Article[];
  stories: Article[];
};

export async function getArticles() {
  try {
    const response = await get<ArticlesResponse>("articles");

    if (!response.ok) {
      return {
        error: response.data?.message || "Failed to fetch articles",
        status: response.status,
      };
    }

    return {
      data: response.data,
      status: response.status,
    };
  } catch (error) {
    console.error("Failed to fetch articles:", error);
    return {
      error: "An unexpected error occurred while fetching articles",
      status: 500,
    };
  }
}
