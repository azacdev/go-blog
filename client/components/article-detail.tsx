import Link from "next/link";
import Image from "next/image";
import { formatDistanceToNow } from "date-fns";

import { Article } from "@/types/articles";
import { Button } from "@/components/ui/button";

interface ArticleDetailProps {
  article: Article | null;
}

export default function ArticleDetail({ article }: ArticleDetailProps) {
  // Format the date to show how long ago it was posted
  const formattedDate =
    article?.CreatedAt &&
    formatDistanceToNow(new Date(article.CreatedAt), {
      addSuffix: true,
    });

  return (
    <div className="py-8 w-full">
      {/* Author and Meta Information */}
      <div className="flex flex-col md:flex-row items-start mb-6">
        <div className="mb-4 md:mb-0">
          <Link href={`/author/${article?.User.ID}`}>
            <div className="relative w-16 h-16 rounded-full overflow-hidden mr-5">
              <Image
                src={
                  article?.User.Image || "/placeholder.svg?height=64&width=64"
                }
                alt={article?.User.Name || "user-image"}
                fill
                className="object-cover"
              />
            </div>
          </Link>
        </div>
        <div className="">
          <div className="flex flex-wrap items-center gap-3 mb-2">
            <Link
              href={`/author/${article?.User.ID}`}
              className="text-lg font-medium hover:underline"
            >
              {article?.User.Name}
            </Link>
            <Button variant="outline" size="sm" className="h-8 rounded-full">
              Follow
            </Button>
          </div>
          <p className="text-muted-foreground mb-2 text-sm">
            {article?.Content}
          </p>
          <div className="flex items-center text-sm text-muted-foreground">
            <span>{formattedDate}</span>
            <span className="mx-2">•</span>
            <span>4 min read</span>
          </div>
        </div>
      </div>

      {/* Article Title */}
      <h1 className="text-3xl md:text-4xl font-bold mb-6">{article?.Title}</h1>

      {/* Featured Image */}
      <div className="relative w-full h-[400px] mb-8 rounded-lg overflow-hidden">
        <Image
          src={article?.Image || "/placeholder.svg?height=400&width=800"}
          alt={article?.Title || "title"}
          fill
          className="object-cover"
          priority
        />
      </div>

      {/* Article Content */}
      <div
        className="prose prose-lg max-w-none mb-8"
        dangerouslySetInnerHTML={{ __html: article?.Content || "" }}
      />
    </div>
  );
}
