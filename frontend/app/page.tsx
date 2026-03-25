"use client";

import { useState, useEffect } from "react";
import Image from "next/image";
import Link from "next/link";

function useIsMobile(breakpoint = 768) {
  const [isMobile, setIsMobile] = useState(false);

  useEffect(() => {
    const mediaQuery = window.matchMedia(`(max-width: ${breakpoint}px)`);

    const handleChange = () => setIsMobile(mediaQuery.matches);

    handleChange();
    mediaQuery.addEventListener("change", handleChange);

    return () => mediaQuery.removeEventListener("change", handleChange);
  }, [breakpoint]);

  return isMobile;
}

export default function Hero() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const isMobile = useIsMobile();

  return (
    <>
      <nav className="fixed top-0 left-0 right-0 z-50 bg-transparent">
        <div className="mx-auto max-w-7xl px-6 sm:px-8">
          <div className="flex h-16 sm:h-20 items-center justify-between">
            <Link href="/" className="flex items-center">
              <Image
                src="/assets/images/Icon arah - General Footage.png"
                alt="Logo"
                width={40}
                height={40}
                priority
              />
            </Link>

            <div className="hidden md:flex items-center gap-5">
              <Link
                href="/auth/login"
                className="text-sm text-white font-semibold hover:text-white">
                Log in
              </Link>
              <Link
                href="/signup"
                // className="rounded-lg bg-blue-600 px-5 py-2 text-sm font-semibold text-white hover:bg-blue-500"
                className="flex items-center gap-4 rounded-md bg-gradient-to-r from-blue-600 to-blue-400 px-5 py-2 text-sm font-semibold text-white shadow-inner shadow-white hover:from-blue-500 hover:to-blue-300 transition-all duration-200">
                Sign Up
              </Link>
            </div>

            <button
              onClick={() => setIsMenuOpen(!isMenuOpen)}
              className="md:hidden text-xl text-white">
              ☰
            </button>
          </div>
        </div>
      </nav>

      <section className="relative flex flex-col h-[120vh] sm:h-[150vh] md:h-[200vh] min-h-screen bg-black overflow-hidden">
        {/* Background */}
        <Image
          src="/assets/images/Background.png"
          alt="Background"
          fill
          priority
          className="object-cover"
        />
        <div className="absolute inset-0 bg-linear-to-b from-transparent via-black/60 to-black" />

        {/* TEXT CONTENT */}
        <div className="relative z-10 mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 pt-28 text-center">
          <h1 className="text-4xl sm:text-6xl md:text-7xl font-bold text-white leading-tight">
            Forex{" "}
            <span className="text-transparent bg-clip-text bg-linear-to-r from-white via-white to-indigo-400">
              Community
            </span>{" "}
            <span className="text-transparent bg-clip-text bg-linear-to-r from-indigo-400 to-blue-600">
              Indonesia
            </span>
          </h1>

          <p className="mt-5 text-base sm:text-4xl text-white/90">
            Belajar Forex Gratis Bareng Komunitas Aktif
          </p>

          <div className="mt-8 flex flex-col sm:flex-row gap-4 justify-center">
            <a className="px-7 py-3 rounded-full border border-white/40 text-white shadow-inner shadow-white">
              Akses Gratis Komunitas
            </a>
            <a className="px-7 py-3 rounded-full border border-white/40 text-white shadow-inner shadow-white">
              Download Resource Gratis
            </a>
          </div>
        </div>

        <div className="absolute bottom-20 sm:-bottom-20 md:left-1/2 md:-bottom-50 md:w-screen md:-translate-x-1/2">
          {/* {!isMobile && (
            <div className="absolute inset-x-0 top-0 h-32 bg-gradient-to-b from-black to-transparent z-10" />
          )} */}

          <Image
            src="/assets/images/Trade Performance.png"
            alt="Trade Performance Dashboard"
            width={1080}
            height={isMobile ? 1080 : 640}
            priority
            className="relative w-full h-auto"
            style={{
              filter: `${isMobile ? "" : "drop-shadow(0 30px 60px rgba(0,0,0,0.6))"}`,
            }}
          />
        </div>

        <div className="relative flex flex-col items-center justify-end flex-1 pb-12">
          <button
            className="
              flex items-center gap-4
              rounded-xl
              bg-gradient-to-r from-blue-600 to-blue-400
              px-5 py-4 text-xl font-semibold text-white
              shadow-inner shadow-white
              hover:from-blue-500 hover:to-blue-300
              transition-all duration-200
            ">
            <span className="relative w-6 h-6">
              <Image
                src="/assets/images/Icon arah - General Footage.png"
                alt="Logo"
                fill
                className="object-contain"
              />
            </span>
            Gabung Komunitas Gratis
          </button>
          <p className="mt-5 text-center font-bold text-lg sm:text-4xl text-white/90 max-w-sm md:max-w-2xl">
            Tempat belajar & berkembang bersama trader Indonesia
          </p>
        </div>
      </section>

      {isMenuOpen && (
        <div className="fixed inset-0 z-40 bg-black/95 backdrop-blur-lg md:hidden">
          <div className="flex flex-col items-center justify-center h-full gap-6">
            <Link
              href="/auth/login"
              className="text-lg text-white"
              onClick={() => setIsMenuOpen(false)}>
              Log in
            </Link>
            <Link
              href="/signup"
              className="rounded-lg bg-blue-600 px-6 py-2.5 text-sm font-semibold text-white"
              onClick={() => setIsMenuOpen(false)}>
              Sign Up
            </Link>
            <button onClick={() => setIsMenuOpen(false)} className="mt-8 text-white/50">
              ✕ Close
            </button>
          </div>
        </div>
      )}
    </>
  );
}
