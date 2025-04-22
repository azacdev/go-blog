import Link from "next/link";
import Image from "next/image";

import { formatDistanceToNow } from "date-fns";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";

interface User {
  id: string;
  name: string;
  image: string;
  description: string;
}

interface Article {
  id: string;
  title: string;
  content: string;
  image: string;
  createdAt: string;
  readTime: number;
  user: User;
  tags: string[];
}

interface ArticleDetailProps {
  article: Article;
}

export default function ArticleDetail({ article }: ArticleDetailProps) {
  // Format the date to show how long ago it was posted
  const formattedDate = formatDistanceToNow(new Date(article.createdAt), {
    addSuffix: true,
  });

  return (
    <div className="max-w-4xl mx-auto px-4">
      <div className="py-8">
        {/* Author and Meta Information */}
        <div className="flex flex-col md:flex-row items-start mb-6">
          <div className="md:w-1/6 mb-4 md:mb-0">
            <Link href={`/author/${article.user.id}`}>
              <div className="relative w-16 h-16 rounded-full overflow-hidden">
                <Image
                  src={
                    article.user.image || "/placeholder.svg?height=64&width=64"
                  }
                  alt={article.user.name}
                  fill
                  className="object-cover"
                />
              </div>
            </Link>
          </div>
          <div className="md:w-5/6">
            <div className="flex flex-wrap items-center gap-3 mb-2">
              <Link
                href={`/author/${article.user.id}`}
                className="text-lg font-medium hover:underline"
              >
                {article.user.name}
              </Link>
              <Button variant="outline" size="sm" className="h-8 rounded-full">
                Follow
              </Button>
            </div>
            <p className="text-muted-foreground mb-2 text-sm">
              {article.user.description}
            </p>
            <div className="flex items-center text-sm text-muted-foreground">
              <span>{formattedDate}</span>
              <span className="mx-2">•</span>
              <span>{article.readTime} min read</span>
            </div>
          </div>
        </div>

        {/* Article Title */}
        <h1 className="text-3xl md:text-4xl font-bold mb-6">{article.title}</h1>

        {/* Featured Image */}
        <div className="relative w-full h-[400px] mb-8 rounded-lg overflow-hidden">
          <Image
            src={article.image || "/placeholder.svg?height=400&width=800"}
            alt={article.title}
            fill
            className="object-cover"
            priority
          />
        </div>

        {/* Article Content */}
        <div
          className="prose prose-lg max-w-none mb-8"
          dangerouslySetInnerHTML={{ __html: article.content }}
        />

        {/* Tags */}
        <div className="border-t pt-6">
          <div className="flex flex-wrap gap-2">
            {article.tags.map((tag, index) => (
              <Link href={`/tags/${tag.toLowerCase()}`} key={index}>
                <Badge variant="secondary" className="px-3 py-1 text-sm">
                  {tag}
                </Badge>
              </Link>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
