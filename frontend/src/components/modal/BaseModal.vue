<!-- src/components/modal/BaseModal.vue -->
<template>
    <div v-if="show" class="fixed inset-0 flex items-center justify-center bg-gray-700 bg-opacity-50 z-50">
        <div class="bg-white p-4 rounded shadow-lg max-w-lg w-full">
            <header class="mb-4">
                <slot name="header">
                    <h3 class="text-lg font-medium">Modal Title</h3>
                </slot>
            </header>
            <div class="mb-4">
                <slot>
                    Modal content goes here.
                </slot>
            </div>
            <footer class="flex justify-end">
                <slot name="footer">
                    <button @click="$emit('close')"
                        class="px-4 py-2 bg-gray-500 text-white rounded hover:bg-gray-600">Close</button>
                </slot>
            </footer>
        </div>
    </div>
</template>

<script>
export default {
    name: 'BaseModal',
    props: {
        show: {
            type: Boolean,
            default: false
        }
    }
}
</script>  
