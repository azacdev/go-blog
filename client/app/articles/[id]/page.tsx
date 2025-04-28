import { getArticleById } from "@/actions/articles";
import ArticleDetail from "@/components/article-detail";

// This would typically come from a database
const mockArticle = {
  id: "1",
  title: "Building Modern Web Applications with Next.js",
  content: `<p>Next.js has revolutionized the way we build React applications. With its powerful features like server-side rendering, static site generation, and API routes, it provides a comprehensive solution for modern web development.</p>
  <p>In this article, we'll explore how to leverage these features to build performant, SEO-friendly applications that provide an excellent user experience.</p>
  <h2>Server Components</h2>
  <p>One of the most exciting features in the latest versions of Next.js is Server Components. These allow you to write components that run only on the server, reducing the JavaScript sent to the client and improving performance.</p>
  <p>Server Components are perfect for data fetching and can significantly improve your application's loading times.</p>`,
  image: "/placeholder.svg?height=400&width=800",
  createdAt: new Date().toISOString(),
  readTime: 6,
  user: {
    id: "user1",
    name: "Sarah Johnson",
    image: "/placeholder.svg?height=64&width=64",
    description:
      "Senior Frontend Developer and creator of various open-source projects. Passionate about creating intuitive user experiences and teaching web development.",
  },
  tags: ["Next.js", "React", "Web Development", "JavaScript"],
};

export default async function ArticlePage({
  params,
}: {
  params: { id: string };
}) {
  const id = await params.id;
  const result = await getArticleById(id);
  const article = result.data?.article;

  return (
    <div className="min-h-screen flex justify-center items-center container mx-auto px-4 pt-20 max-w-5xl">
      <ArticleDetail article={article || null} />
    </div>
  );
}
