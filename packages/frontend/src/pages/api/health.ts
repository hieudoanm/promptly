import type { NextApiRequest, NextApiResponse } from 'next';

type ResponseData = {
	status: string;
};

const handler = (
	request: NextApiRequest,
	response: NextApiResponse<ResponseData>,
) => {
	const { method } = request;
	if (method !== 'GET') {
		response.setHeader('Allow', ['GET']);
		return response.status(405).end(`Method ${method} Not Allowed`);
	}
	return response.status(200).json({ status: 'ok' });
};

export default handler;
