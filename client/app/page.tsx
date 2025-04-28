import Image from "next/image";
import Link from "next/link";

import { Footer } from "@/components/footer";
import { Header } from "@/components/header";
import { getArticles } from "@/actions/articles";

export default async function Home() {
  const result = await getArticles();

  const articles = result.data;

  return (
    <div className="bg-white">
      <div className="container mx-auto px-4 pt-20">
        {/* Header would be imported here */}
        <Header />

        {/* Featured Posts Section */}
        <section className="featured-posts mb-12">
          <div className="section-title mb-6">
            <h2 className="text-2xl font-bold">
              <span className="border-b-2 border-gray-800 pb-1">Featured</span>
            </h2>
          </div>
          {!result.error && articles && articles.featured.length > 0 ? (
            <div className="grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3">
              {articles.featured.map((article) => (
                <div
                  key={article.ID}
                  className="card bg-white rounded-lg shadow-md overflow-hidden"
                >
                  <div className="flex flex-col md:flex-row">
                    <div className="md:w-5/12">
                      <Link href={`/articles/${article.ID}`}>
                        <div
                          className="thumbnail h-48 md:h-full bg-cover bg-center"
                          style={{ backgroundImage: `url(${article.Image})` }}
                        ></div>
                      </Link>
                    </div>
                    <div className="md:w-7/12 p-4">
                      <div className="card-block">
                        <h2 className="card-title text-xl font-bold mb-2">
                          <Link
                            href={`/articles/${article.ID}`}
                            className="text-gray-900 hover:text-gray-700"
                          >
                            {article.Title}
                          </Link>
                        </h2>
                        <h4 className="card-text text-gray-600 mb-4 line-clamp-3">
                          {article.Content}
                        </h4>
                        <div className="metafooter border-t pt-3">
                          <div className="wrapfooter flex items-center">
                            <span className="meta-footer-thumb mr-2">
                              <Link href="#">
                                <Image
                                  className="author-thumb rounded-full"
                                  src={article.User.Image || "/placeholder.svg"}
                                  alt={article.User.Name}
                                  width={30}
                                  height={30}
                                />
                              </Link>
                            </span>
                            <span className="author-meta">
                              <span className="post-name">
                                <Link
                                  href="#"
                                  className="text-gray-900 font-medium"
                                >
                                  {article.User.Name}
                                </Link>
                              </span>
                              <br />
                              <span className="post-date text-sm text-gray-500">
                                {article.CreatedAt}
                              </span>
                              <span className="dot mx-1">•</span>
                              <span className="post-read text-sm text-gray-500">
                                6 min read
                              </span>
                            </span>
                            <span className="post-read-more ml-auto">
                              <Link
                                href="#"
                                title="Read Story"
                                className="text-gray-500 hover:text-gray-700"
                              >
                                <svg
                                  className="svgIcon-use"
                                  width="25"
                                  height="25"
                                  viewBox="0 0 25 25"
                                >
                                  <path
                                    d="M19 6c0-1.1-.9-2-2-2H8c-1.1 0-2 .9-2 2v14.66h.012c.01.103.045.204.12.285a.5.5 0 0 0 .706.03L12.5 16.85l5.662 4.126a.508.508 0 0 0 .708-.03.5.5 0 0 0 .118-.285H19V6zm-6.838 9.97L7 19.636V6c0-.55.45-1 1-1h9c.55 0 1 .45 1 1v13.637l-5.162-3.668a.49.49 0 0 0-.676 0z"
                                    fillRule="evenodd"
                                  ></path>
                                </svg>
                              </Link>
                            </span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="flex items-center justify-center">
              <p className="text-gray-600 text-center w-full">
                No featured available.
              </p>
            </div>
          )}
        </section>

        {/* All Stories Section */}
        <section className="recent-posts mb-12">
          <div className="section-title mb-6">
            <h2 className="text-2xl font-bold">
              <span className="border-b-2 border-gray-800 pb-1">
                All Stories
              </span>
            </h2>
          </div>

          {!result.error && articles && articles.featured.length > 0 ? (
            <div className="grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3">
              {articles?.featured.map((story) => (
                <div
                  key={story.ID}
                  className="card bg-white rounded-lg shadow-md overflow-hidden"
                >
                  <Link href={`/articles/${story.ID}`}>
                    <Image
                      className="w-full h-48 object-cover"
                      src={story.Image || "/placeholder.svg"}
                      alt={story.Title}
                      width={600}
                      height={400}
                    />
                  </Link>
                  <div className="card-block p-4">
                    <h2 className="card-title text-xl font-bold mb-2">
                      <Link
                        href={`/articles/${story.ID}`}
                        className="text-gray-900 hover:text-gray-700"
                      >
                        {story.Title}
                      </Link>
                    </h2>
                    <h4 className="card-text text-gray-600 mb-4 line-clamp-3">
                      {story.Content}
                    </h4>
                    <div className="metafooter border-t pt-3">
                      <div className="wrapfooter flex items-center">
                        <span className="meta-footer-thumb mr-2">
                          <Link href="#">
                            <Image
                              className="author-thumb rounded-full"
                              src={story.User.Image || "/placeholder.svg"}
                              alt={story.User.Name}
                              width={30}
                              height={30}
                            />
                          </Link>
                        </span>
                        <span className="author-meta">
                          <span className="post-name">
                            <Link
                              href={`/articles/${story.ID}`}
                              className="text-gray-900 font-medium"
                            >
                              {story.User.Name}
                            </Link>
                          </span>
                          <br />
                          <span className="post-date text-sm text-gray-500">
                            {story.CreatedAt}
                          </span>
                          <span className="dot mx-1">•</span>
                          <span className="post-read text-sm text-gray-500">
                            6 min read
                          </span>
                        </span>
                        <span className="post-read-more ml-auto">
                          <Link
                            href="#"
                            title="Read Story"
                            className="text-gray-500 hover:text-gray-700"
                          >
                            <svg
                              className="svgIcon-use"
                              width="25"
                              height="25"
                              viewBox="0 0 25 25"
                            >
                              <path
                                d="M19 6c0-1.1-.9-2-2-2H8c-1.1 0-2 .9-2 2v14.66h.012c.01.103.045.204.12.285a.5.5 0 0 0 .706.03L12.5 16.85l5.662 4.126a.508.508 0 0 0 .708-.03.5.5 0 0 0 .118-.285H19V6zm-6.838 9.97L7 19.636V6c0-.55.45-1 1-1h9c.55 0 1 .45 1 1v13.637l-5.162-3.668a.49.49 0 0 0-.676 0z"
                                fillRule="evenodd"
                              ></path>
                            </svg>
                          </Link>
                        </span>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="flex items-center justify-center">
              <p className="text-gray-600 text-center w-full">
                No stories available.
              </p>
            </div>
          )}
        </section>

        <Footer />
      </div>
    </div>
  );
}
