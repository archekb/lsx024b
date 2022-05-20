<template>
    <v-card flat class="client-settings-content">
        <v-toolbar color="second">
            <v-btn icon="mdi-close" @click="close" class="mr-2" />
            <v-toolbar-title>Client Settings</v-toolbar-title>
            <v-tooltip anchor="start">
                <template v-slot:activator="{ props }">
                    <v-btn icon="mdi-restore" v-bind="props" @click="reset" />
                </template>
                <span>Reset</span>
            </v-tooltip>
        </v-toolbar>

        <div class="client-settings-content">
            <v-row no-gutters justify="center">
                <v-col v-for="setting, settingKey in settings" :key="settingKey" cols="12" sm="6">

                    <v-list v-if="settingKey=='color'" lines="two" subheader>
                        <v-list-subheader>{{ nameFormatter(settingKey) }}</v-list-subheader>
                        <v-list-item v-for="prop, propKey in setting" :key="propKey" :title="nameFormatter(propKey)" :subtitle="prop?.description">
                            <template v-slot:append>
                                <v-list-item-avatar start class="ml-3">
                                    <ColorSelector :color="prop?.value" :onColor="c => set(`${settingKey}.${propKey}.value`, c)"/>
                                </v-list-item-avatar>
                            </template>
                        </v-list-item>
                    </v-list>

                    <v-list v-if="['voltage', 'current', 'power', 'temperature'].includes(settingKey)" lines="two" subheader>
                        <v-list-subheader>{{ nameFormatter(settingKey) }}</v-list-subheader>
                        <v-list-item v-for="prop, propKey in setting" :key="propKey" :subtitle="prop?.description" style="flex-direction: column; align-items:flex-start">
                            <div class="my-4" style="width: 100%;">
                                <v-text-field type="number" :model-value="prop.from" :label="nameFormatter(propKey) + ' From'" hide-details variant="underlined" full-width @update:modelValue="c => set(`${settingKey}.${propKey}.from`, c)">
                                    <template v-slot:append>
                                        {{ prop?.unit }}
                                    </template>
                                </v-text-field>
                            </div>
                            <div class="my-4" style="width: 100%;">
                                <v-text-field type="number" :model-value="prop.to" :label="nameFormatter(propKey) + ' To'" hide-details variant="underlined" full-width @update:modelValue="c => set(`${settingKey}.${propKey}.to`, c)">
                                    <template v-slot:append>
                                        {{ prop?.unit }}
                                    </template>
                                </v-text-field>
                            </div>
                        </v-list-item>
                    </v-list>

                </v-col>
            </v-row>
        </div>
    </v-card>
</template>

<script>
import ColorSelector from "@/components/ColorSelector"
import { nameFormatter } from "@/utils"

export default {
    name: 'web-client-settings',
    components: {
        ColorSelector
    },
    props: {
        settings: Object,
        close: Function,
        set: Function,
        reset: Function
    },
    methods: {
        nameFormatter,
    }
}
</script>

<style>
.client-settings-content {
    width: 100%;
    height: 100%;
}
</style>