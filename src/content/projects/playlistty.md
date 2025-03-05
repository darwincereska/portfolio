---
title: 'Playlistty'
description: 'A command-line tool for transferring playlists between music streaming services'
pubDate: 'Feb 25 2025'
heroImage: '/projects/playlistty.png'
source: 'https://github.com/darwincereska/playlistty'
---
**Playlistty** is a project still in development, waiting for more music services to open back up their api.

#### Why?
I created Playlistty to solve a common problem I faced - **the need to move playlists between different music streaming services**.

As someone who uses both **Spotify** and **YouTube Music**, I was frustrated by the lack of easy and **free** ways to transfer my carefully curated playlists between platforms.

#### Supported platforms.
The current supported platforms supported are:
- Spotify
- YouTube Music

#### Future plans.
My future plans are to integrate more streaming services, but due to their apis not being available right now, I just have to wait.
Future platforms:
- Apple Music
- Deezer
- Sound Cloud
- Tidal

#### How it works?
Playlistty is a simple command-line tool that makes this process seamless.

Built with Go, it securely connects to the streaming services' APIs and handles the transfer of playlists while maintaining all the song metadata.

I focused on making it both powerful and user-friendly, with features like OAuth authentication, playlist management, and cross-platform song searching.

The tool supports basic operations like creating new playlists and clearing existing ones, but the main functionality is the ability to transfer playlists between services.

It handles the OAuth flow for secure API access and stores credentials safely in a config file.

#### The future of the project.
Since releasing it, I've been happy to see others finding it useful for managing their music libraries across platforms.

The project is open source under the **MIT** license and I welcome contributions from the community.