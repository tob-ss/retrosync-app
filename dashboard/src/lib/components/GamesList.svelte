<script lang="ts">
   // import { GetSaves } from "../../wailsjs/go/main/App";
   // import { GetHeaders } from "../../wailsjs/go/main/App";
   // import { Svroller } from "svrollbar";
    import SaveFile from "./SaveFile.svelte";
    import { getSaves } from "$lib/scripts/getSaves";
    import { getHeaders } from "$lib/scripts/getHeaders";
    import LoadingSaves from "./LoadingSaves.svelte";

    let data = $state([]);
    
    /*getSaves().then(data => {
        console.log(data);
    }); /*

    /*(async () => {
        const data = await getSaves();
        
    }); */

    let headers = $state([]);

    /*getHeaders(data).then(headers => {
        console.log(headers);
    }); */

    /*(async () => {
        const headers = await getHeaders(data);

        
    }); */

    let test_thumbnail = "https://f003.backblazeb2.com/file/retrosync-thumbnails/Sony+-+PlayStation+Portable/Named_Boxarts/Dragon+Ball+Z+-+Tenkaichi+Tag+Team+(Europe)+(En%2CFr%2CDe%2CEs%2CIt).png"
    
    
</script>


<div class="h-full w-full overflow-auto">
    
        {#await getSaves()}
            <div class="h-full">
                <div class="h-12  py-24 flex justify-center">
                    <h1 class="text-4xl font-bold">
                        <p class="skeleton rounded-xl">Value for Item 1</p>
                    </h1>
                </div>
                <div class="flex justify-center flex-wrap">
                <div>
                    <LoadingSaves />
                </div>
                <div>
                    <LoadingSaves />
                </div>
                <div>
                    <LoadingSaves />
                </div>
                <div>
                    <LoadingSaves />
                </div>
                </div>
            </div>
        {:then data}
        {#await getHeaders(data)}
            <p>Getting Headers...</p>
        {:then headers}
        <div class="flex flex-col text-[#D7D6FC] font-heebo">
        {#each headers as header}
            <div class="h-full">
                <div class="h-12  py-24 flex justify-center">
                    <h1 class="text-4xl font-bold">{header}</h1>
                </div>
                <div class="flex justify-center flex-wrap">
                {#each data as save}
                <div>
                    {#if header === save.Header}
                        <SaveFile
                            gameName={save.Name}
                            console={save.Console}
                            device={save.Device}
                            timeMod={save.TimeString}
                            thumbnail={save.Thumbnail}
                            />
                    {/if}
                </div>
                {/each}
                </div>
            </div>
        {/each}
        </div>
        {/await}
        {/await}
</div>

<style>
.skeleton {
    animation: skeleton-loading 1s linear infinite alternate;
    color: transparent;
}

@keyframes skeleton-loading {
    0% {
        background-color: #2C2F48; /* FROM Color 1 */
    }
    100% {
        background-color: #E2E2E2; /* TO Color 2 */
    }
} 
</style>