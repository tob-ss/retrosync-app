<script lang="ts">
    
    import Sidebar from '$lib/components/Sidebar.svelte';

    function beforeunload(event: BeforeUnloadEvent) {
		event.preventDefault();
		return event.returnValue = '';
	}

    let { children } = $props();
    
</script>

<svelte:window on:beforeunload={beforeunload}/>


<div class="h-full w-full grid grid-cols-5 grid-rows-5 gap-4">
    <div class="row-span-5">
        <Sidebar />
    </div>
    <div class="col-span-4 row-span-5">
        {@render children()}
        
    </div>

</div>