import React, { useEffect, useState } from 'react';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import { map } from 'lodash';

import { getValueFromWidgetData, setKeyValueToState, renderEditorFieldComponents, updateLocalisationValue } from '../../helper'
import { typography } from '../../styles/globalStyles'
import EditorContainer from '../EditorContainer'

const useStyles = makeStyles(() => ({
    widgetContainer: {
        marginTop: "1rem",
        marginBottom: "1rem",
    },
    telcoDetailItemContainer: {
        margin: "auto",
        paddingRight: "1.5rem"
    },
    telcoDetailTitle: {
        ...typography.heading1,
        margin: "auto"
    },
    telcoDetailLabel: {
        ...typography.body1,
    },
    telcoDetailIcon: {
        maxWidth: "1.5rem",
        paddingRight: "0.25rem",
        float: "right"
    },
    telcoDetailTitleImage: {
        maxWidth: "1.1rem",
    }
}));

const GridMaxFourComponent = ({widget, updateEditorComponent, resetSelectedCard, shouldRenderEditorComponent, pageId, fetchPageData}) => {
    const [previewComponents, setPreviewComponents] = useState(null);
    const [widgetData, setWidgetData] = useState({});
    const [editing, setEditing] = useState(false);
    const [saving, setSaving] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const classes = useStyles();

    // If widget is selected for editing, render EditorContainer
    useEffect(() => {
        if (shouldRenderEditorComponent && widgetData) {
            // updateLocalisationValueFunc is a wrapper function for updateLocalisationValue to avoid having to pass in params(that are local to this scope) to renderEditorFieldComponents function.
            const updateLocalisationValueFunc = (localeKey, localeValue) => {
                updateLocalisationValue({localeKey, localeValue, pageId, setErrorMessage, setSaving, setEditing, fetchPageData})
            }
            // Generate EditorFieldComponents based on widgetData
            const editorFieldComponents = renderEditorFieldComponents({widgetData, saving, setWidgetData, editing, setEditing, initState, updateLocalisationValueFunc})
            // Invoke updateEditorComponent to render the EditorComponent
            updateEditorComponent(
                <EditorContainer 
                    errorMessage={errorMessage} 
                    editorFieldComponents={editorFieldComponents}
                    resetFunction={resetSelectedCard}
                />
            )
        }
    }, [shouldRenderEditorComponent, widgetData, saving, errorMessage])

    // Render Preview component
    useEffect(() => {
        setPreviewComponents(
            map(widgetData, (value, type) => {
                if(type === "telco-detail") {
                    const parsedImage = getValueFromWidgetData(widgetData, type, "data.image")
                    const parsedSubtitle = getValueFromWidgetData(widgetData, type, "data.subtitle")
                    const parsedTitleText = getValueFromWidgetData(widgetData, type, "data.title")
                    const parsedTitleImage = getValueFromWidgetData(widgetData, type, "data.titleImage")
                    const parsedTitle = parsedTitleImage ? <img src={parsedTitleImage} alt="tooltip" className={clsx(classes.telcoDetailTitleImage)} /> : parsedTitleText
                    return (<Grid item xs key={parsedSubtitle}>
                        <Box>
                            <Grid container direction="row" className={clsx(classes.telcoDetailItemContainer)}>
                                <Grid item xs={6}><img src={parsedImage} alt="icon" className={clsx(classes.telcoDetailIcon)} /></Grid>
                                <Grid item xs={6} className={clsx(classes.telcoDetailTitle)}>{parsedTitle}</Grid>
                                <Grid item xs={6}></Grid>
                                <Grid item xs={6} className={clsx(classes.telcoDetailLabel)}>{parsedSubtitle}</Grid>
                            </Grid>
                        </Box>
                    </Grid>)
                }
            })
        )
    }, [widgetData])

    // initState initializes the widgetData state for use to render the preview as well as the EditorContainer
    const initState = () => {
        if (!widget) return
        switch(widget.type) {
            case "telco-detail" : {
                setKeyValueToState(widget, "telco-detail", "data.image", setWidgetData)
                setKeyValueToState(widget, "telco-detail", "data.subtitle", setWidgetData)
                setKeyValueToState(widget, "telco-detail", "data.title", setWidgetData)
                setKeyValueToState(widget, "telco-detail", "data.titleImage", setWidgetData)
            }
        }
    }

    useEffect(() => {
        initState()
    }, [widget])

    return (previewComponents && <>
        {map(previewComponents, previewComponent => previewComponent)}
    </>)
}

export default GridMaxFourComponent;