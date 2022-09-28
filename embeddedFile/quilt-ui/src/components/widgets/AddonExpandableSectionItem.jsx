import React, { useEffect, useState } from 'react';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import { map } from 'lodash';

import { getValueFromWidgetData, setKeyValueToState, renderEditorFieldComponents, updateLocalisationValue } from '../../helper'
import { typography, colors, commonStyles } from '../../styles/globalStyles'
import EditorContainer from '../EditorContainer'

const useStyles = makeStyles(() => ({
    addOnCard: {
        background: colors.general01Color,
        borderRadius: "10px",
        margin: "0.35rem",
        padding: "0.5rem 0.3rem",
        minHeight: "3.5rem",
        ...commonStyles.shadow

    },
    addOnCardColor: {
        width: "4px",
        height: "2rem",
        display: "inline-block",
        marginTop: "50%",
        borderRadius: "3px",
    },
    addOnTitle: {
        ...typography.heading3,
        lineHeight: "1",
        marginBottom: "0.2rem"
    },
    addOnDescription: {
        ...typography.body2,
        lineHeight: "1",
        color: colors.text01Color
    },
    addOnSubTtile: {
        ...typography.body2,
        color: colors.primary01Color100,
        float: "right",
        marginBottom: "0.2rem"
    },
    addOnImage: {
        maxWidth: "0.8rem",
        maxHeight: "0.8rem",
        float: "right"
    },
}));

const AddonExpandableSectionItem = ({widget, updateEditorComponent, resetSelectedItemCard, shouldRenderEditorComponent, pageId, fetchPageData}) => {
    const [previewComponents, setPreviewComponents] = useState([]);
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
                    resetFunction={resetSelectedItemCard}
                />
            )
        }
    }, [shouldRenderEditorComponent, widgetData, saving, errorMessage])

    // Render Preview component
    useEffect(() => {
        setPreviewComponents(
            map(widgetData, (value, type) => {
                if(type === "addon-tile") {                    
                    const parsedColor = getValueFromWidgetData(widgetData, "addon-tile", "data.color")
                    const parsedTitle = getValueFromWidgetData(widgetData, "addon-tile", "data.title")
                    const parsedDescription = getValueFromWidgetData(widgetData, "addon-tile", "data.description")
                    const parsedSubtitle = getValueFromWidgetData(widgetData, "addon-tile", "data.subtitle")
                    const parsedImage = getValueFromWidgetData(widgetData, "addon-tile", "data.image")
                    const parsedId = getValueFromWidgetData(widgetData, "addon-tile", "data.id")
                    return (<Grid item xs={6} key={parsedId}>
                        <Box className={clsx(classes.addOnCard)}>
                            <Grid container direction="row">
                                <Grid item xs={1}>
                                    <Box
                                        className={clsx(classes.addOnCardColor)} 
                                        style={{backgroundColor:parsedColor}}
                                    ></Box>
                                </Grid>
                                <Grid item xs={7}>
                                    <Typography 
                                        variant="body1"
                                        className={clsx(classes.addOnTitle)}
                                    >
                                        {parsedTitle}
                                    </Typography>
                                    <Typography 
                                        variant="body1"
                                        className={clsx(classes.addOnDescription)}
                                    >
                                        {parsedDescription}
                                    </Typography>
                                </Grid>
                                <Grid item xs={4}>
                                    <Grid container direction="column">
                                        <Grid item>
                                            <Typography 
                                                variant="body1"
                                                className={clsx(classes.addOnSubTtile)}
                                            >
                                                {parsedSubtitle}
                                            </Typography>
                                        </Grid>
                                        <Grid item>
                                            <img src={parsedImage} alt="icon" className={clsx(classes.addOnImage)} />
                                        </Grid>

                                    </Grid>
                                </Grid>
                            </Grid>
                        </Box>
                    </Grid>)
                } else {
                    return <div>{`Unsupported Component: ${type}`}</div>
                }
            })
        )
    }, [widgetData])

    // initState initializes the widgetData state for use to render the preview as well as the EditorContainer
    const initState = () => {
        if (!widget) return

        if (widget.type === "addon-tile") {
            setKeyValueToState(widget, "addon-tile", "data.color", setWidgetData)
            setKeyValueToState(widget, "addon-tile", "data.title", setWidgetData)
            setKeyValueToState(widget, "addon-tile", "data.description", setWidgetData)
            setKeyValueToState(widget, "addon-tile", "data.subtitle", setWidgetData)
            setKeyValueToState(widget, "addon-tile", "data.image", setWidgetData)
            setKeyValueToState(widget, "addon-tile", "data.id", setWidgetData)
        }
    }

    useEffect(() => {
        initState()
    }, [widget])

    return (previewComponents && <>
        {map(previewComponents, previewComponent => previewComponent)}
    </>)
}

export default AddonExpandableSectionItem;