<template>
    <button :class="buttonClasses" @click="handleClick">
        <slot></slot>
    </button>
</template>

<script>
export default {
    name: 'BaseButton',
    props: {
        type: {
            type: String,
            default: 'button'
        },
        buttonClasses: {
            type: String,
            default: 'px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600'
        }
    },
    methods: {
        handleClick(event) {
            this.$emit('click', event)
        }
    }
}
</script>