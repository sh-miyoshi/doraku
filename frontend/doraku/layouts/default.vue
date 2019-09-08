<template>
  <div>
    <header>
      <nuxt-link to="/">
        <img src="~/assets/logo.png" width="150" />
      </nuxt-link>
    </header>

    <nuxt />

    <footer>
      <nuxt-link to="/">Topページへ</nuxt-link>&emsp;&emsp;
      <nuxt-link to="/about" no-prefetch>Dorakuとは？</nuxt-link>&emsp;&emsp;
      <a href="https://github.com/sh-miyoshi/doraku/issues/new/choose"
        >Contact Support</a
      >
    </footer>
  </div>
</template>

<style>
body {
  background-image: url('~assets/background.png');
  background-repeat: no-repeat;
  background-color: #dee9a2;
  background-size: cover;
  text-align: center;
  margin: 0;
  padding: 0;
  font-size: 18px;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

header {
  text-align: left;
  margin: 30px;
}

footer {
  margin: 30px;
  font-size: 14px;
}
</style>
