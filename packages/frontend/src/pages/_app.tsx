import '@chat/styles/globals.css';
import type { AppProps } from 'next/app';
import Head from 'next/head';
import { Geist, Geist_Mono } from 'next/font/google';
import { FC } from 'react';

const geistSans = Geist({
	variable: '--font-geist-sans',
	subsets: ['latin'],
});

const geistMono = Geist_Mono({
	variable: '--font-geist-mono',
	subsets: ['latin'],
});

const App: FC<AppProps> = ({ Component, pageProps }) => {
	return (
		<>
			<Head>
				<title>Chat</title>
			</Head>
			<div className={`${geistSans.className} ${geistMono.className}`}>
				<Component {...pageProps} />
			</div>
		</>
	);
};

export default App;
