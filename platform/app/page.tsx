"use client"
import Lottie from 'react-lottie-player'
import animationJson from "../public/animation.json"
import Link from 'next/link';

export default function Home() {

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <section className="flex flex-col justify-between gap-6 sm:gap-10 md:gap-16 lg:flex-row">
        <div className="flex flex-col justify-center sm:text-center lg:py-12 lg:text-left xl:w-5/12 xl:py-24">
          <p className="mb-4 font-semibold text-indigo-500 md:mb-6 md:text-lg xl:text-xl">Trading with AI</p>

          <h1 className="mb-8 text-4xl font-bold text-black sm:text-5xl md:mb-12 md:text-6xl">Revolutionary way to trade. Start writting your promt about your portfolio and learn the best trade advantages</h1>

          <p className="mb-8 leading-relaxed text-gray-500 md:mb-12 lg:w-4/5 xl:text-lg">Only for Canto chain</p>

          <div className="flex flex-col gap-2.5 sm:flex-row sm:justify-center lg:justify-start">
            <Link href="/invest" className="inline-block rounded-lg bg-indigo-500 px-8 py-3 text-center text-sm font-semibold text-white outline-none ring-indigo-300 transition duration-100 hover:bg-indigo-600 focus-visible:ring active:bg-indigo-700 md:text-base">Let's Start</Link>

          </div>
        </div>

        <div className="h-48 overflow-hidden w-1/2 rounded-lg bg-gray-100 shadow-lg lg:h-auto xl:w-5/12">
          <Lottie
            animationData={animationJson}
            play
            loop
            speed={0.4}
          />        </div>
      </section>
    </main>
  );
}
