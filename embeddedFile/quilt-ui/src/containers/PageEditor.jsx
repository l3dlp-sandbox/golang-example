import React, { useState, useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Card from '@material-ui/core/Card';
import Typography from '@material-ui/core/Typography';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import { get, map } from 'lodash';

import UserBaseCard from '../components/widgets/UserBaseCard';
import GridMaxFour from '../components/widgets/GridMaxFour';
import AddonExpandable from '../components/widgets/AddonExpandable';
import GridMaxTwoExpandable from '../components/widgets/GridMaxTwoExpandable';
import PageNavigationBar from '../components/widgets/PageNavigationBar';
import List from '../components/widgets/List';
import BorderedList from '../components/widgets/BorderedList';
import RequestService from '../requestService';
import { typography, colors } from '../styles/globalStyles'

// This is the list of supported widgets. Update this list as we progressively add more widgets
const SUPPORTED_WIDGETS = {
    "user-base-card": UserBaseCard,
    "grid-max-four": GridMaxFour,
    "addon-expandable": AddonExpandable,
    "grid-max-two-expandable": GridMaxTwoExpandable,
    "list": List,
    "page-navigation-bar": PageNavigationBar,
    "bordered-list": BorderedList,
}

const useStyles = makeStyles(() => ({
    mainContainer: {
        marginTop: "2rem",
    },
    widgetContainer: {
        marginBottom: '1rem',
        padding: '1.5rem',
    },
    selectField: {
        width: "100%",
    },
    selectContainer: {
        margin: "1rem 0rem",
    },
    phoneContainer: {
        position: "sticky",
        top: "4rem",
        width: "375px", // iphone 6 size
        height: "667px", // iphone 6 size
        border: "1px solid blue",
        overflow: "scroll",
        backgroundColor: colors.background01Color
    },
    blurred: {
        opacity: 0.2,
    },
    clickable: {
        cursor: "pointer",
    },
    header: {
        ...typography.editorHeading1,
        margin: "1rem 0rem"
    },
    editOrderButton: {
        color: colors.primary01Color100,
        opacity: 0.2,
        ...typography.actionText,
        margin: "1.7rem 0rem",
        float: "right"
    }
}));

// The PageEditor component presents a preview of the selected page. It also renders an editor for the user to edit localisation values
const PageEditor = ({pageId, selectedLocale, country}) => {

    // widgetsWithIDs state is needed to assign a unique ID to each widget in the pageData. At point of writing, the quilt team is in the midst of adding IDs to widgets but this has not been added to all widgets hence I opted to dynamically set a ID in code for now.
    const [widgetsWithIDs, setWidgetsWithIDs] = useState(null);
    // currentHoveredWidget state is used to track the widget that user's mouse is hovering on so that we can blur out the other widgets
    const [currentHoveredWidget, setCurrentHoveredWidget] = useState(null);
    // selectedWidgetToEdit state is used to track the widget that the user is editing. It is used to determine which widget has the right to update the editorComponent
    const [selectedWidgetToEdit, setSelectedWidgetToEdit] = useState(null);
    // editorComponent state is used to render the editor where user's can update the localisation values. It is updated by the selectedWidgetToEdit
    const [editorComponent, setEditorComponent] = useState(null);

    const history = useHistory();

    // Dynamically assign a unique ID to each widget based on it's index
    const mapWidgetsWithID = (pageData) => {
        if (!pageData || get(pageData, "widgets.length") === 0) return null
        const result = []
        for(let i=0; i<pageData.widgets.length; i += 1) {
            const widget = pageData.widgets[i]
            result.push({
                ...widget,
                uiID: `${widget.type}-${i}`
            })
        }
        return result
    }

    // Fetch page data, invokes the get page data api to retrieve the page data from quilt's api server
    const fetchPageData = async () => {
        if (selectedLocale && country) {
            const url = process.env.GET_PAGE_DATA_PATH + pageId // This is to aid local development to use an alternate server instead of quilt where auth is needed with each request. 
            let response
            try {
                response = await RequestService.get(url, {
                    headers: {
                        "Accept-Language": `${selectedLocale}-${country}`
                    }
                })
            } catch (error) {
                // If authToken is no longer valid, redirect user to login page
                if (error.response.status === 401) {
                    history.push(`/`)
                }
            }
            // Set widgets with response from fetchPageData
            setWidgetsWithIDs(mapWidgetsWithID(response.data))
        }
    }

    // When page loads fetch page data
    useEffect(() => {
        fetchPageData()
    }, [pageId, selectedLocale])

    // updateEditorComponent is passed to widget components for to update the editorComponent
    const updateEditorComponent = (newEditorCompoent) => {
        setEditorComponent(newEditorCompoent)
    }

    // resetEditorComponent is passed to widget components to hide the editorComponent and display the selectWidgetToEditSection instead.
    const resetEditorComponent = () => {
        setSelectedWidgetToEdit(null)
    }

    const classes = useStyles();

    // renderWidget checks if the widget is supported.
    // If widget is unsupported, it renders a placeholder
    // If widget is supported, it renders the widget component
    // the function also blurs the widget if users hovers over the widget in the selectWidgetToEditSection
    const renderWidget = (widget) => {
        if (Object.keys(SUPPORTED_WIDGETS).includes(widget.type)) {
            // Blur all other widgets except the currentHoveredWidget
            const shouldBlur = currentHoveredWidget && (currentHoveredWidget !== widget.uiID)
            // Only allow the selected widget to be able to update editorComponent
            const shouldRenderEditorComponent = selectedWidgetToEdit === widget.uiID
            // Dynamically set supported widget
            const WidgetComponent = SUPPORTED_WIDGETS[widget.type]
            return <Box
                key={widget.uiID} 
                className={shouldBlur ? clsx(classes.blurred) : null}
            >
                <WidgetComponent 
                    widget={widget}
                    pageId={pageId}
                    updateEditorComponent={updateEditorComponent}
                    resetEditorComponent={resetEditorComponent}
                    shouldRenderEditorComponent={shouldRenderEditorComponent}
                    fetchPageData={fetchPageData}
                />
            </Box>
        } 
        return <div key={widget.uiID}>{`Unsupported Widget: ${widget.type}`}</div>
    }

    // previewSection renders the preview of the page based on the supported widgets
    const previewSection = () => (
        <Box className={clsx(classes.phoneContainer)}>
        {widgetsWithIDs && map(widgetsWithIDs, (widget) => (
            renderWidget(widget)
        ))}
        </Box>
    )
    
    // selectWidgetToEditSection enables the user to select the widget that they want to update
    const selectWidgetToEditSection = () => (
        <Grid item xs={6}>
            <Grid container direction="row">
                <Grid item xs={9}>
                    <Typography 
                        className={clsx(classes.header)}
                    >
                        Select a section to edit
                    </Typography>
                </Grid>
                <Grid item xs={3}>
                    <Typography 
                        className={clsx(classes.editOrderButton)}
                    >
                        Edit Order
                    </Typography>
                </Grid>

            </Grid>
            {widgetsWithIDs && map(widgetsWithIDs, (widget) => {
                const supportedWidgets = Object.keys(SUPPORTED_WIDGETS)
                if (!supportedWidgets.includes(widget.type)) return null
                // eslint-disable-next-line
                return <Box key={widget.uiID}>
                    <Card 
                        onMouseEnter={() => {
                            setCurrentHoveredWidget(widget.uiID)
                        }}
                        onMouseLeave={() => {
                            setCurrentHoveredWidget(null)
                        }}
                        onClick={()=>{
                            setSelectedWidgetToEdit(widget.uiID)
                        }}
                        className={clsx(classes.widgetContainer, classes.clickable)}
                    >
                        <Typography 
                            className={clsx(classes.header)}
                        >
                            {widget.type}
                        </Typography>
                    </Card>
                </Box>
            })}
        </Grid>
    )

    return (<Grid 
        container 
        direction="row"
        className={clsx(classes.mainContainer)}
    >
        <Grid item sm={12} md={6}>
            {previewSection()}
        </Grid>
        {selectedWidgetToEdit &&
            <Grid item sm={12} md={6}>
                {editorComponent}
            </Grid>
        }
        {!selectedWidgetToEdit && selectWidgetToEditSection()}
    </Grid>);
};

export default PageEditor;
