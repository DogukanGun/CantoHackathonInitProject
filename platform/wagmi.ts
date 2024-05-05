import { http, createConfig } from 'wagmi'
import { canto } from 'wagmi/chains'

export const config = createConfig({
  chains: [canto],
  transports: {
    [canto.id]: http()
  },
})