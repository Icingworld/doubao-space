async function request(path, options = {}) {
  const response = await fetch(path, {
    headers: {
      Accept: 'application/json',
      ...options.headers,
    },
    ...options,
  })

  if (!response.ok) {
    throw new Error(`请求失败：${response.status}`)
  }
  return response.json()
}

export function getHealth() {
  return request('/api/health')
}
