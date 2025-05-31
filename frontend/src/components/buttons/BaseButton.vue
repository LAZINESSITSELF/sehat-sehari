<template>
    <button :class="buttonClasses" @click="$emit('click', $event)">
      <slot></slot>
    </button>
  </template>
  
  <script>
  export default {
    name: "BaseButton",
    props: {
      buttonClasses: {
        type: String,
        default: 'px-4 py-2 bg-blue-500 text-white rounded'
      },
      type: {
        type: String,
        default: 'button'
      }
    }
  }
  </script>