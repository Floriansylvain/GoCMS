<script setup lang="ts">
const emits = defineEmits(['close'])
const props = defineProps<{
    title?: string,
    description: string,
    type: 'error' | 'success'
}>()

</script>

<template>
    <div class="modal">
        <header class="error-bg" v-if="type === 'error'">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round"
                    d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126zM12 15.75h.007v.008H12v-.008z" />
            </svg>
            <h2>{{ props.title ?? 'Échec !' }}</h2>
        </header>
        <header class="success-bg" v-else>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round"
                    d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <h2>{{ props.title ?? 'Succès !' }}</h2>
        </header>
        <main>
            <p>{{ props.description }}</p>
            <button class="button-primary" @click="$emit('close')">Fermer</button>
        </main>
    </div>
</template>

<style scoped>
.modal {
    position: absolute;
    top: 50%;
    left: 50%;

    box-shadow: #0002 0 0 15px;
    border-radius: var(--radius);
    overflow: hidden;

    max-width: 100%;
    width: 500px;
    transform: translate(-50%, -50%);
}

.success-bg {
    background-color: var(--primary-light);
}

.error-bg {
    background-color: var(--secondary-light);
}

header {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 8px;

    color: var(--neutral-dark);
}

main {
    display: flex;
    flex-direction: column;
    gap: 16px;
    padding: 16px;
    background-color: white;
}

main p {
    margin: 0;
}

main button {
    align-self: flex-end;
}

svg {
    width: 50px;
    color: var(--neutral-dark);
}
</style>