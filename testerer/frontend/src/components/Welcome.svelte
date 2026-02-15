<script lang="ts">
    import { fade } from 'svelte/transition'
    import Setup from './Setup.svelte';

    let visible = $state(false);

    let showSetup = $state(false);

    setTimeout(() => {
        visible = true
    })

    setTimeout(() => {
            visible = false
            setTimeout(() => {
                showSetup = true
            }, 1050)
        }, 4000)
    


</script>

{#if visible}
    <h1 transition:fade={{ duration: 1000 }} class="relative flex items-center text-5xl font-medium text-[#D7D6FC] font-heebo">Welcome to RetroSync</h1>
{/if}

{#if showSetup}
    <Setup />
{/if}

<style>
   
</style>