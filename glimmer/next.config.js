module.exports = {
  distDir: 'build',
  future: {
    webpack5: true,
  },
  images: {
    domains: [
      'raw.githubusercontent.com',
    ],
  },
  async headers() {
    return [
      {
        source: '/_next/image',
        headers: [
          {
            key: 'Cache-Control',
            value: 'public,max-age=31536000,immutable',
          },
        ],
      },
    ]
  },
};
