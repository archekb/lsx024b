<template>
    <v-card flat class="mx-auto" max-width="1920">

        <v-toolbar color="primary" dark extended flat>
            <v-toolbar-title>{{ this.title }}</v-toolbar-title>
            <v-btn icon="mdi-cog" @click="showClientSettings = !showClientSettings" :active="showClientSettings" /> 
        </v-toolbar>

        <v-card class="mx-auto" max-width="1020" style="margin-top: -50px;" v-if="!showClientSettings">
            <v-toolbar flat>
                <v-toolbar-title class="grey--text">
                    <v-icon size="small" :color="connected" icon="mdi-record" />
                    {{ nameFormatter(tab) }}
                </v-toolbar-title>
                <v-btn :icon="modern ? 'mdi-format-list-bulleted' : 'mdi-apps'" @click="modern = !modern" class="mr-2" /> 
                <v-menu anchor="start">
                    <template v-slot:activator="{ props }">
                        <v-btn icon="mdi-dots-vertical" v-bind="props" />
                    </template>

                    <v-list class="grey lighten-3">
                        <v-list-item v-for="item, key in items" :key="key" :disabled="key == tab" @click="tab = key">
                            {{ nameFormatter(key) }}
                        </v-list-item>
                    </v-list>
                </v-menu>
            </v-toolbar>

            <v-divider></v-divider>

            <v-window v-model="tab">
                <v-window-item v-for="item, key in items" :key="key" :value="key" style="min-height: 200px; margin: 10px;">
                    <v-card flat>
                        <ControllerRealTime v-if="key == 'real_time' && modern" :data="item" :settings="clientSettings.settings"/>
                        <v-table v-else-if="key == tab" fixed-header>
                            <thead>
                                <tr>
                                    <th class="text-left">Name</th>
                                    <th class="text-left">Value</th>
                                </tr>
                            </thead>
                            <tbody>
                            <tr v-for="[k, v]  in Object.entries(item)" :key="k">
                                <td>{{ nameFormatter(k) }}</td>
                                <td :style="{ color: colorFormatter(clientSettings.settings, k, v, '#OOOOOO') }">{{ v }} {{ unitFormatter(k, v) }}</td>
                            </tr>
                            </tbody>
                        </v-table>
                    </v-card>
                </v-window-item>
            </v-window>
        </v-card>

        <v-card class="mx-auto" max-width="1020" style="margin-top: -50px;" v-else>
            <ClientSettings :settings="clientSettings.settings" :set="clientSettings.Set" :reset="clientSettings.Reset" :close="() => showClientSettings = false"  />
        </v-card>

    </v-card>

</template>

<script>
import ClientSettings from "@/components/ClientSettings";
import ControllerRealTime from "@/components/ControllerRealTime";
import { useControllerStateStore } from "@/store/controller_state";
import { useClientSettingsStore } from "@/store/client_settings";
import { storeToRefs } from 'pinia';
import { onUnmounted } from 'vue';
import { nameFormatter, unitFormatter, colorFormatter } from "@/utils";

export default {
    name: 'home-page',
    components: {
        ClientSettings,
        ControllerRealTime,
    },
    setup() {
        const clientSettings = useClientSettingsStore();
        const stateStore = useControllerStateStore();
        const { state, inProgress, error } = storeToRefs(stateStore);

        stateStore.start();
        onUnmounted(stateStore.stop);
        return {
            clientSettings,
            state,
            inProgress,
            error
        }
    },
    data() {
        return {
            tab: "real_time",
            modern: true,
            showClientSettings: false
        }
    },
    methods: {
        unitFormatter,
        nameFormatter,
        colorFormatter,
    },
    computed: {
        items() {
            return this.state?.device || {}
        },
        title() {
            const title = `${this.state?.model || 'Epever (Epsolar)'} Controller Monitor`;
            if (this.state?.model && document.title != title) {
                document.title = title;
            }
            return title
        },
        connected() {
            return this.state?.connected ? '#689F38' : '#D32F2F'
        }
    },
 }
</script>
