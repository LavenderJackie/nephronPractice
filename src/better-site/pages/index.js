import Head from 'next/head'
import {useState, useEffect, useRef} from 'react'

export default function Home() {
  return (
    <div className="flex flex-col items-center justify-center py-2">
      <div className="h-40" />
      <div className="flex flex-row items-center justify-center">
        <div className="bg-green-500 rounded-xl flex h-60 w-full px-20">
        </div>
      </div>
    </div>
  )
}
