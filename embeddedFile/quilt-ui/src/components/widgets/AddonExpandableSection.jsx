import React, { useEffect, useState } from 'react';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import Card from '@material-ui/core/Card';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import { get, map } from 'lodash';

import AddonExpandableSectionItem from './AddonExpandableSectionItem';
import { parseQUIPrefix, getValueFromWidgetData, getValueFromParsedObject, setKeyValueToState, renderEditorFieldComponents, updateLocalisationValue } from '../../helper'
import { typography } from '../../styles/globalStyles'
import EditorContainer from '../EditorContainer'

const useStyles = makeStyles(() => ({
    sectionHeader: {
        ...typography.heading2,
        paddingLeft: "0.35rem",
    },
    legendTitle: {
        ...typography.body2,
        display: "inline",
    },
    legendColor: {
        width: "10px",
        height: "3px",
        display: "inline-block",
        marginBottom: "0.1rem",
        borderRadius: "3px",
    },
    clickableCard: {
        cursor: "pointer",
        width: "100%",
        margin: "0.5rem",
        padding: "1rem",
        textAlign: "center"
    }
}));

const AddonExpandableSection = ({widget, updateEditorComponent, resetSelectedSectionCard, shouldRenderEditorComponent, pageId, fetchPageData}) => {
    const [childComponents, setChildComponents] = useState(null);
    const [legends, setLegends] = useState(null);
    const [widgetData, setWidgetData] = useState({});
    const [editing, setEditing] = useState(false);
    const [saving, setSaving] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const [clickableCards, setClickableCards] = useState([]);
    // selectedCard state is used to track which AddonExpandableSectionItem component is being selected
    const [selectedCard, setSelectedCard] = useState(null)
    const classes = useStyles();

    // If widget is selected for editing, render EditorContainer
    useEffect(() => {
        // Note that if a selectedCard is set, this AddonExpandableSection component will not update the editorComponent since the selected AddonExpandableSectionItem will be the one updating it instead.
        if (shouldRenderEditorComponent && !selectedCard && widgetData) {
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
                    resetFunction={resetSelectedSectionCard}
                    clickableCards={clickableCards}
                />
            )
        }
    }, [shouldRenderEditorComponent, widgetData, saving, errorMessage, selectedCard])

    // resetSelectedCard is passed as a prop to AddonExpandableSectionItem to unselect itself.
    const resetSelectedItemCard = () => {
        setSelectedCard(null)
    }

    // initState initializes the widgetData state for use to render the preview as well as the EditorContainer
    // It also instantiates the clickableCards state that is used to enable the user to select the card that they want to edit
    const initState = () => {
        if (!widget) return
        setChildComponents(null)
        setClickableCards([])
        setKeyValueToState(widget, "title", "data.sectionHeader.title", setWidgetData)
        const parsedSectionHeaderLegendsCount = get(widget, "data.sectionHeader.legends.length", 0)
    
        if (parsedSectionHeaderLegendsCount > 0) {
            const selectedLegends = get(widget, "data.sectionHeader.legends")
            for(let i=0; i<parsedSectionHeaderLegendsCount; i++) {
                const legend = selectedLegends[i]
                setKeyValueToState(legend, `legend-${i}.title`, "data.title", setWidgetData)
                setKeyValueToState(legend, `legend-${i}.color`, "data.color", setWidgetData)
            }
        }
        setChildComponents(map(widget.data.items, item => {
            // Get cardTitle for displaying in ClickableCard
            const {parsedKeyValue} = parseQUIPrefix(get(item, "data.title"))
            const cardTitle = getValueFromParsedObject(parsedKeyValue)
            const clickableCard = (<Card
                key={cardTitle}
                className={clsx(classes.clickableCard)}
                onClick={()=>{
                    setSelectedCard(cardTitle)
                }}
            >
                {cardTitle}
            </Card>)
            setClickableCards(prevState => [
                ...prevState,
                clickableCard
            ])

            const shouldRenderItemEditorComponent = selectedCard === cardTitle

            return <AddonExpandableSectionItem
                key={cardTitle}
                widget={item}
                updateEditorComponent={updateEditorComponent}
                shouldRenderEditorComponent={shouldRenderItemEditorComponent}
                resetSelectedItemCard={resetSelectedItemCard}
                pageId={pageId}
                fetchPageData={fetchPageData}
            />
        }))
    }

    // This effect is used to re-generate the setClickableCards after a card has been clicked
    useEffect(() => {
        initState()
    }, [selectedCard])

    useEffect(() => {
        initState()
    }, [widget])

    // Render legend
    useEffect(() => {
        const types = Object.keys(widgetData)
        // Check if section has legend
        if (types.includes("legend-0.title")) {
            const legend0Title = getValueFromWidgetData(widgetData, "legend-0.title", "data.title")
            const legend0Color = getValueFromWidgetData(widgetData, "legend-0.color", "data.color")
            const legend1Title = getValueFromWidgetData(widgetData, "legend-1.title", "data.title")
            const legend1Color = getValueFromWidgetData(widgetData, "legend-1.color", "data.color")
            setLegends(<>
                <Grid key={legend0Color} item xs={6}>
                    <Box className={clsx(classes.legendColor)} style={{backgroundColor:legend0Color}} />
                    <Typography className={clsx(classes.legendTitle)}>{legend0Title}</Typography>
                </Grid>
                <Grid key={legend1Color} item xs={6}>
                    <Box className={clsx(classes.legendColor)} style={{backgroundColor:legend1Color}} />
                    <Typography className={clsx(classes.legendTitle)}>{legend1Title}</Typography>
                </Grid>
            </>)
        }
    }, [widgetData])

    const parsedSectionHeaderTitle = getValueFromWidgetData(widgetData, "title", "data.sectionHeader.title")
    return (widgetData && <Grid item xs key={parsedSectionHeaderTitle}>
        <Box>
            <Grid container direction="row">
                <Grid item xs={3} className={clsx(classes.sectionHeader)}>{parsedSectionHeaderTitle}</Grid>
                <Grid item xs={4} ></Grid>
                <Grid item xs={5}>
                    <Grid container direction="row">
                        {legends}
                    </Grid>
                </Grid>
                <Grid container direction="row">
                    { map(childComponents, childComponent => childComponent)}
                </Grid>
            </Grid>
        </Box>
    </Grid>)
}

export default AddonExpandableSection;