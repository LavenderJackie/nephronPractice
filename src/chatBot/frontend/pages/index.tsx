import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'

import axios from 'axios'
import {useState, useEffect, useRef} from 'react'

const Home: NextPage = () => {
  const [npcs, setNPCS] = useState(null)
  const [messages, setMessages] = useState([])
  const [typed, setTyped] = useState("")

  useEffect(() => {
    axios.get("http://localhost:8080/npcs/tavern keep").then((res) => {
      setNPCS(res.data)
    })
  }, [messages])

  const npcRespond = () => {
    for (let i = npcs.responses.length - 1; i >= 0; i--) {
      if(messages[messages.length - 1].includes(npcs.responses[i].prompt)) {
        messages.push(npcs.name + ": " + npcs.responses[i].phrase)
        break
      }
    }
  }

  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <div className={`block rounded-xl flex flex-col mx-16 bg-green-600`}>
          <div className={`block rounded-xl m-4 w-80 h-8 p-2 bg-green-100 overflow-y-scroll ${styles.outputLog}`}>
            {messages.map(e => {return <p className={`${e.includes("you: ") ? "text-black" : "text-blue-700"} font-bold`}> {e} </p>})}
          </div>
          <input className="block rounded-xl m-4 mt-0 w-80 h-8 p-2 bg-white border border-border transparent focus:outline-none"
                  placeholder="type here to chat"
                  value={typed}
                  onChange={(e) => {setTyped(e.target.value)}}/>
          <button className="block rounded-xl m-4 mt-0 w-80 p-2 bg-green-300 hover:bg-green-400 text-center font-extrabold"
                  onClick={() => {
                    messages.push("you: " + typed)
                    npcRespond()
                    setTyped("")
                  }}> submit </button>
        </div>
      </main>
    </div>
  )
}

export default Home
