<script lang="ts">
    import { GetSaves } from "../../wailsjs/go/main/App";
    import { GetHeaders } from "../../wailsjs/go/main/App";
    import { Svroller } from "svrollbar";
    import SaveFile from "./SaveFile.svelte";

  //  let data: Array<Record<string, any>> = GetSaves();

    GetSaves().then(data => {
        console.log(data);
    });

    let data = $state([]);

    (async () => {
        const data = await GetSaves();
        
    });

    let headers = $state([]);

    GetHeaders(data).then(headers => {
        console.log(headers);
    });

    (async () => {
        const headers = await GetHeaders(data);

        
    });
    

    let test_thumbnail = "https://f003.backblazeb2.com/file/retrosync-thumbnails/Sony+-+PlayStation+Portable/Named_Boxarts/Dragon+Ball+Z+-+Tenkaichi+Tag+Team+(Europe)+(En%2CFr%2CDe%2CEs%2CIt).png"
    
    
</script>


<div class="h-full w-full overflow-y-scroll">
    
        {#await GetSaves()}
            <p>Getting Saves...</p>
        {:then data}
        {#await GetHeaders(data)}
            <p>Getting Headers...</p>
        {:then headers}
        <div class="flex flex-col text-[#D7D6FC] font-heebo">
        {#each headers as header}
            <div class="h-full text-center">
                <div class="h-12  py-24">
                    <h1 class="text-4xl font-bold">{header}</h1>
                </div>
                <div class="flex justify-center flex-wrap">
                {#each data as save}
                <div>
                    {#if header === save.Date_String}
                        <SaveFile 
                            gameName={save.Game_Name}
                            console={save.Console}
                            device={save.Device}
                            timeMod={save.Time_Modified}
                            thumbnail={test_thumbnail}
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