"use client";

import Link from "next/link";
import Image from "next/image";
import { useState } from "react";
import { Search } from "lucide-react";
import { useAtom } from "jotai/react";
import { userAtom } from "@/store/user-store";

export default function Navbar() {
  const [user] = useAtom(userAtom);
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  const handleLogout = () => {
    console.log("Logging out...");
  };

  return (
    <nav className="bg-white fixed top-0 left-0 right-0 z-10 shadow-sm">
      <div className="container mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          {/* Logo */}
          <Link href="/" className="flex-shrink-0 font-bold text-xl">
            Go Blog
          </Link>

          {/* Mobile menu button */}
          <div className="md:hidden">
            <button
              type="button"
              className="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-gray-500"
              aria-expanded="false"
              onClick={toggleMenu}
            >
              <span className="sr-only">Open main menu</span>
              <svg
                className="block h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                aria-hidden="true"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M4 6h16M4 12h16M4 18h16"
                />
              </svg>
            </button>
          </div>

          {/* Desktop menu */}
          <div className="hidden md:flex md:items-center md:justify-between md:flex-1">
            <div className="flex-1"></div>
            <div className="flex items-center">
              {/* Search form */}
              <form className="relative mr-4">
                <input
                  className="form-input rounded-md py-2 px-4 border border-gray-300 bg-white placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-gray-500 focus:border-gray-500 sm:text-sm"
                  type="text"
                  placeholder="Search"
                />
                <span className="absolute right-3 top-2 text-gray-400">
                  <Search size={20} />
                </span>
              </form>

              {/* User menu */}
              {!user || user.ID === 0 ? (
                <Link
                  href="/login"
                  className="ml-4 px-3 py-2 rounded-md text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50"
                >
                  Login
                </Link>
              ) : (
                <div className="ml-4 relative">
                  <div>
                    <button
                      type="button"
                      className="flex items-center max-w-xs rounded-full text-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500"
                      id="user-menu-button"
                      aria-expanded="false"
                      aria-haspopup="true"
                      onClick={toggleMenu}
                    >
                      <span className="sr-only">Open user menu</span>
                      <Image
                        className="h-8 w-8 rounded-full"
                        src={
                          user.Image || "/placeholder.svg?height=32&width=32"
                        }
                        alt={user.Name}
                        width={32}
                        height={32}
                      />
                      <span className="ml-2 text-gray-700">{user.Name}</span>
                    </button>
                  </div>

                  {/* Dropdown menu */}
                  {isMenuOpen && (
                    <div
                      className="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none"
                      role="menu"
                      aria-orientation="vertical"
                      aria-labelledby="user-menu-button"
                      tabIndex={-1}
                    >
                      <Link
                        href="/articles/create"
                        className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                        role="menuitem"
                      >
                        Create article
                      </Link>
                      <button
                        onClick={handleLogout}
                        className="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                        role="menuitem"
                      >
                        Logout
                      </button>
                    </div>
                  )}
                </div>
              )}
            </div>
          </div>
        </div>
      </div>

      {/* Mobile menu, show/hide based on menu state */}
      {isMenuOpen && (
        <div className="md:hidden">
          <div className="px-2 pt-2 pb-3 space-y-1 sm:px-3">
            {!user || user.ID === 0 ? (
              <Link
                href="/login"
                className="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50"
              >
                Login
              </Link>
            ) : (
              <>
                <div className="px-3 py-2 flex items-center">
                  <Image
                    className="h-8 w-8 rounded-full mr-2"
                    src={user.Image || "/placeholder.svg?height=32&width=32"}
                    alt={user.Name}
                    width={32}
                    height={32}
                  />
                  <span className="text-gray-700">{user.Name}</span>
                </div>
                <Link
                  href="/articles/create"
                  className="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50"
                >
                  Create article
                </Link>
                <button
                  onClick={handleLogout}
                  className="block w-full text-left px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50"
                >
                  Logout
                </button>
              </>
            )}
          </div>
          <div className="pt-4 pb-3 border-t border-gray-200">
            <form className="px-4">
              <div className="relative">
                <input
                  className="form-input w-full rounded-md py-2 px-4 border border-gray-300 bg-white placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-gray-500 focus:border-gray-500 sm:text-sm"
                  type="text"
                  placeholder="Search"
                />
                <span className="absolute right-3 top-2 text-gray-400">
                  <Search size={20} />
                </span>
              </div>
            </form>
          </div>
        </div>
      )}
    </nav>
  );
}
